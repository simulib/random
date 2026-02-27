// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
F random variate generator.

This file implements sampling from the F distribution using the fact
that an F-distributed random variable can be expressed as the ratio
of two independent chi-squared random variables.

Reference:

C. Forbes, M. Evans, N. Hastings, B. Peacock.
Statistical Distributions.
Wiley, 2011 (4th ed.), p. 108.
*/

package random

import "fmt"

// GetFSampler returns a sampler function for the F distribution.
//
// Distribution: F
//
// Parameters:
//   - d1: degree of freedom (>0)
//   - d2: degree of freedom d2 (> 0)
//
// The probability density function is:
//
//	f(x) = (d1/d2)^(d1/2) * x^(d1/2 - 1) / B(d1/2, d2/2) *
//	       (1 + d1*x/d2)^(-(d1+d2)/2),  for x ≥ 0
//	     = 0,                           otherwise
//
// where B(y, z) = Gamma(y) * Gamma(z) / Gamma(y + z) is
// the Beta function, and Gamma() is the Gamma function.
//
// Example:
//
//	sampler, _ := r.GetFSampler(1.0, 2.0)
//	x := sampler()
func (rng *Random) GetFSampler(d1 float64, d2 float64) (func() float64, error) {
	if d1 <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: d1 must be positive")
	}
	if d2 <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: d2 must be positive")
	}
	return func() float64 {
		g1 := rng.Gamma(d1 / 2.0)
		g2 := rng.Gamma(d2 / 2.0)
		return (g1 * d2) / (g2 * d1)
	}, nil
}
