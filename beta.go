// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Beta random variate generator.

This file implements sampling from the Beta distribution using the
"Beta variates via gamma variates" method described by Devroye.

Reference:

[1] L. Devroye,
    Non-Uniform Random Variate Generation,
    Springer, 1986, p. 432.
*/

package random

import "fmt"

// GetBetaSampler returns a sampler function for the Beta distribution.
//
// Distribution: Beta
//
// Parameters:
//   - nu: shape parameter (>0)
//   - omega: shape parameter (>0)
//
// The probability density function is
//
//	f(x) = x^(nu-1) * (1-x)^(omega-1) / B(nu, omega),  for 0 ≤ x ≤ 1
//		 = 0,                                          otherwise
//
// where B(nu, omega) = Gamma(nu) * Gamma(omega) / Gamma(nu + omega) is
// the Beta function, and Gamma() is the Gamma function.
//
// Example:
//
//	sampler, _ := r.GetBetaSampler(1.0, 2.0)
//	x := sampler()
func (rng *Random) GetBetaSampler(nu float64, omega float64) (func() float64, error) {
	if nu <= 0 {
		return nil, fmt.Errorf("invalid parameter: nu must be positive")
	}
	if omega <= 0 {
		return nil, fmt.Errorf("invalid parameter: omega must be positive")
	}
	return func() float64 {
		x := rng.Gamma(nu)
		y := rng.Gamma(omega)
		return x / (x + y)
	}, nil
}
