// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Bernoulli random variate generator

This file implements the inversion method for generating Bernoulli distributed
random variates.
*/

package random

import "fmt"

// GetBernoulliSampler returns a sampler function for the Bernoulli distribution.
//
// Distribution: Bernoulli
//
// Parameters:
//   - p: success probability (0 ≤ p ≤ 1).
//
// The probability mass function is
//
//	f(x) = 1 - p,  for x = 0
//	     = p,      for x = 1
//	     = 0,      otherwise
//
// Example:
//
//	sampler, _ := r.GetBernoulliSampler(0.3)
//	x := sampler()
func (rng *Random) GetBernoulliSampler(p float64) (func() int64, error) {
	if p <= 0 || p > 1 {
		return nil, fmt.Errorf("invalid parameter: p must be in the interval [0,1]")
	}
	return func() int64 {
		if rng.Uniform() < p {
			return 1.0
		}
		return 0
	}, nil
}
