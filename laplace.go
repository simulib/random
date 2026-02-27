// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Laplace random variate generator.

This file implements sampling from the Laplace distribution by
generating an exponential random variable and adding or subtracting it
from the location parameter with equal probability.

This construction exploits the symmetry of the Laplace distribution
and provides an efficient sampling method.
*/

package random

import "fmt"

// GetLaplaceSampler returns a sampler for the Laplace distribution.
//
// Distribution: Laplace(a,b) (double exponential)
//
// Parameters:
//   - a: location parameter
//   - b: scale parameter (> 0)
//
// PDF:
//
//	f(x) = (1 / (2 b)) exp(-|x − a| / b),  for −∞ < x < ∞
//
// Example:
//
//	sampler, _ := r.GetLaplaceSampler(0.0, 1.0)
//	x := sampler()
func (rng *Random) GetLaplaceSampler(a float64, b float64) (func() float64, error) {
	if b < 0.0 {
		return nil, fmt.Errorf("invalid parameter: b must be non-negative")
	}
	return func() float64 {
		if rng.Uniform() < 0.5 {
			return a + rng.Exponential()*b
		}
		return a - rng.Exponential()*b
	}, nil
}
