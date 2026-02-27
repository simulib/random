// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Weibull random variate generator.

This file implements sampling from the Weibull distribution
using the inversion method. For efficiency, −log(U) is replaced by
an exponential random variable.

*/

package random

import "fmt"
import "math"

// GetWeibullSampler returns a sampler function for the Weibull distribution.
//
// Distribution: Weibull
//
// Parameters:
//   - lambda: scale (>0)
//   - k: shape (>0)
//
// The density function is
//
//	f(x) = (k/lambda) (x/lambda)^(k-1) exp(-(x/lambda)^k),   for x ≥ 0
//	     = 0,                                 otherwise
//
// Example usage:
//
//	sampler, _ := r.GetWeibullSampler(1.0, 2.0)
//	x := sampler()
func (rng *Random) GetWeibullSampler(lambda float64, k float64) (func() float64, error) {
	if lambda <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: lambda must be positive")
	}
	if k <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: k must be positive")
	}
	return func() float64 { return lambda * math.Pow(rng.Exponential(), 1.0/k) }, nil
}
