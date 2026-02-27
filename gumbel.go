// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Extreme Value (Gumbel) random variate generator.

This file implements sampling from the Gumbel distribution using the
inversion method. For efficiency, the -log(uniform) computation is
replaced by an exponential random variate.

*/

package random

import "fmt"
import "math"

// GetGumbelSampler returns a sampler function for the Gumbel distribution.
//
// Distribution: Gumbel (Extreme value distribution)
//
// Parameters:
//   - a: location parameter
//   - b: scale parameter (>0)
//
// The probability density function is:
//
//	f(x) = exp(-z - exp(-(x-a)/b)) / b,  for -∞ < x < ∞
//
// Example:
//
//	sampler, _ := r.GetGumbelSampler(1.0, 2.0)
//	x := sampler()
func (rng *Random) GetGumbelSampler(a float64, b float64) (func() float64, error) {
	if b <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: theta must be positive")
	}
	return func() float64 { return a - b*math.Log(rng.Exponential()) }, nil
}
