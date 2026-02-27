// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Erlang random variate generator.

This file implements sampling from the Erlang distribution. An
Erlang random variable with k states and rate lambda is equivalent to
a gamma random variable with shape parameter k and scale parameter 1/lambda.
*/

package random

import "fmt"

// GetErlangSampler returns a sampler function for the Erlang distribution.
//
// Distribution: Erlang
//
// Parameters:
//   - lambda: rate parameter (> 0)
//   - k: number of stages k (> 0, integer).
//
// The probability density function is:
//
//	f(x) = lambda^k / (k-1)! * x^(k-1) * exp(-lambda * x),  for x ≥ 0
//	     = 0,                                               otherwise
//
// Example:
//
//	sampler, _ := r.GetErlangSampler(2.0, 3)
//	x := sampler()
func (rng *Random) GetErlangSampler(lambda float64, k int64) (func() float64, error) {
	if lambda <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: lambda must be positive")
	}
	if k <= 0 {
		return nil, fmt.Errorf("invalid parameter: k must be positive")
	}
	return func() float64 { return rng.Gamma(float64(k)) / lambda }, nil
}
