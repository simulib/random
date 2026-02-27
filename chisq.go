// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Chi-squared random variate generator.

This file implements sampling from the chi-squared distribution. A
chi-squared random variable with ν degrees of freedom is equivalent to
a gamma random variable with shape parameter ν/2 and scale parameter 2.
*/

package random

import "fmt"

// GetChiSqSampler returns a sampler function for the chi-squared distribution.
//
// Distribution: Chi squared
//
// Parameters:
//   - nu: shape parameter (degrees of freedom, nu > 0).
//
// The probability density function is:
//
//	f(x) = 2^(-nu/2) / Gamma(nu/2) * x^(nu/2 - 1) * e^(-x/2),  for 0 ≤ x < ∞
//	     = 0,                                               otherwise
//
// where Gamma() is the Gamma function.
//
// Example:
//
//	sampler, _ := r.GetChiSqSampler(3.0)
//	x := sampler()
func (rng *Random) GetChiSqSampler(nu float64) (func() float64, error) {
	if nu <= 0 {
		return nil, fmt.Errorf("invalid parameter: nu must be non-negative")
	}
	return func() float64 { return 2.0 * rng.Gamma(nu/2.0) }, nil
}
