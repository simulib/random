// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Rayleigh random variate generator.

This file implements sampling from the Rayleigh distribution using
the inversion method. If U ~ Uniform(0,1), then X = b * sqrt(-2 * log(U))
follows a Rayleigh(b) distribution. In the implementation, an
equivalent exponential random variate is used instead of -log(U).

Reference:

C. Forbes, M. Evans, N. Hastings, B. Peacock.
Statistical Distributions.
Wiley, 2011 (4th ed.), Chapter 39.
*/

package random

import "fmt"
import "math"

// GetRayleighSampler returns a sampler function for the Rayleigh distribution.
//
// Distribution: Rayleigh
//
// Parameters:
//   - b: scale parameter (>0)
//
// The probability density function is:
//
//	f(x) = x / b^2 * exp(-x^2 / (2 * b^2)),  for 0 <= x < ∞
//	     = 0,                                otherwise
//
// Example:
//
//	sampler, _ := r.GetRayleighSampler(2.0)
//	x := sampler()
func (rng *Random) GetRayleighSampler(b float64) (func() float64, error) {
	if b <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: a must be positive")
	}
	return func() float64 { return b * math.Sqrt(2.0*rng.Exponential()) }, nil
}
