// Copyright (c) 2025 Dieter Fiems
// SPDX-License-Identifier: MIT

// Package random provides high-performance random variate generation
// for stochastic simulation.
//
// The package separates two concerns:
//
//  1. Core random number generators (RNGs), which produce uniform
//     random bits.
//  2. Distribution samplers (Normal, Exponential, Gamma, etc.),
//     which transform uniform randomness into non-uniform variates.
//
// A Random value represents a compiled distribution kernel built on
// top of a specific RNG. The choice of RNG affects performance and
// reproducibility, but does not affect the semantics of the
// distributions.
//
// Typical usage:
//
//	rng  := mt64.New(seed)
//	r    := rand.NewRand(rng)
//	x    := r.Normal()
//	y    := r.Exponential()
//
// A Random is not safe for concurrent use by multiple goroutines.
// Parallel simulations should use independent Random instances.
//
// Independent Random instances can be created either by constructing
// separate RNGs by using RNG-specific jump-ahead or stream-splitting
// methods.
package random

import (
	"fmt"
	"strings"
)

type Rng64 interface {
	Uint64() uint64
}

// Random provides fast sampling from common probability distributions
// using a single underlying random number generator.
//
// A Random value encapsulates a specific RNG and provides sampling from
// some common distributions. After construction, all sampling methods
// operate without interface dispatch and without knowledge of the
// concrete RNG type.
//
// Random does not reseed or reset the underlying RNG. All samples
// drawn from a Random instance belong to the same random stream.
//
// A Random is not safe for concurrent use by multiple goroutines.
// Create independent Random values for parallel simulations by
// RNG splitting or jump-ahead methods.
type Random struct {
	// Uint64 returns a uniformly distributed 64-bit pseudorandom number
	Uint64 func() uint64
}

// New constructs a Random value backed by the provided RNG.
//
// New uses generics to specialize the distribution algorithms
// to the concrete RNG type at construction time. This allows the
// compiler to inline RNG operations and eliminate interface
// dispatch in the hot sampling paths.
//
// Example:
//
//	rng := mt.New(1234)
//	r   := random.New(rng)
//	x   := r.Normal()
func New[R Rng64](rng R) *Random {

	u64 := func() uint64 { return rng.Uint64() }

	return &Random{
		Uint64: u64,
	}
}

// NewContinuousVariate parses a textual specification of a discrete
// probability distribution and returns a sampling function.
//
// The specification must be of the form:
//
//	DistributionName(param1, param2, ...)
//
// For example:
//
//	"Binomial(3,0.2)"
//
// The returned function generates independent random samples drawn
// from the specified distribution each time it is called.
func (rng *Random) NewContinuousVariate(str string) (func() float64, error) {
	clean := strings.Join(strings.Fields(str), "")
	parts := strings.FieldsFunc(clean, func(r rune) bool {
		return r == '(' || r == ')' || r == ','
	})
	name := strings.ToUpper(parts[0])
	vars := parts[1:]
	val, ok := ContinuousRegistry[name]
	if !ok {
		return nil, fmt.Errorf("no continuous variate for %s", parts[0])
	}
	return val.Sampler(rng, vars)
}

// NewDiscreteVariate parses a textual specification of a discrete
// probability distribution and returns a sampling function.
//
// The specification must be of the form:
//
//	DistributionName(param1, param2, ...)
//
// For example:
//
//	"Binomial(3,0.2)"
//
// The returned function generates independent random samples drawn
// from the specified distribution each time it is called.
func (rng *Random) NewDiscreteVariate(str string) (func() int64, error) {
	clean := strings.Join(strings.Fields(str), "")
	parts := strings.FieldsFunc(clean, func(r rune) bool {
		return r == '(' || r == ')' || r == ','
	})
	name := strings.ToUpper(parts[0])
	vars := parts[1:]
	val, ok := DiscreteRegistry[name]
	if !ok {
		return nil, fmt.Errorf("no discrete variate for %s", parts[0])
	}
	return val.Sampler(rng, vars)
}

// TODO:
//    - implement a sampler for non-central chi-square random variates
//    - implement a sampler for the Dirichlet distribution
//    - implement the alias method
//    - implement sampling for the empirical distribution
//    - implement error distribution
//	  - implement hypergeometric distribution
