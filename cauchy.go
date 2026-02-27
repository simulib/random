// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Cauchy random variate generator

This file implements the inversion method for generating Cauchy distributed
random variates.
*/

package random

import "fmt"
import "math"

// GetCauchySampler returns a sampler function for the Cauchy distribution.
//
// Distribution: Cauchy
//
// Parameters:
//   - a: location parameter (the median)
//   - b: scale parameter (> 0).
//
// The probability density function is:
//
//	f(x) = 1 / (pi * b) / (1 + ((x - a)/b)^2),  for -∞ < x < ∞
//
// Example:
//
//	sampler, _ := r.GetCauchySampler(3.0, 2.0)
//	x := sampler()
func (rng *Random) GetCauchySampler(a, b float64) (func() float64, error) {
	if b <= 0 {
		return nil, fmt.Errorf("invalid parameter: gamma must be positive")
	}
	return func() float64 {
		return a + b*math.Tan(math.Pi*rng.Uniform())
	}, nil
}
