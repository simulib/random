// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Power random variate generator.

This file implements sampling from the Power distribution using
the inversion method. If U ~ Uniform(0,1), then X = b*U^(1/c)
follows a Power(b,c) distribution on [0,b].

Reference:

C. Forbes, M. Evans, N. Hastings, B. Peacock.
Statistical Distributions.
Wiley, 2011 (4th ed.), Chapter 36.
*/

package random

import "fmt"
import "math"

// GetPowerSampler returns a sampler function for the Power distribution.
//
// Distribution: Power
//
// Parameters:
//   - b: scale parameter b (>0)
//   - c: shape parameter c (>0)
//
// The probability density function is:
//
//	f(x) = c * x^(c-1)/b^c,  for 0 <= x <= b
//	     = 0,                otherwise
//
// Example:
//
//	sampler, _ := r.GetPowerSampler(2.0,3.0)
//	x := sampler()
func (rng *Random) GetPowerSampler(b, c float64) (func() float64, error) {
	if b <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: b must be positive")
	}
	if c <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: c must be positive")
	}
	return func() float64 { return b * math.Pow(rng.Uniform(), 1.0/c) }, nil
}
