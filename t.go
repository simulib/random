// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Student-t random variate generator.

This file implements sampling from the Student-t distribution using
the representation 	T = Z / sqrt(G/nu) where Z ~ Normal(0,1) and
G ~ Gamma(nu/2, 2).

Reference:

L. Devroye.
Non-Uniform Random Variate Generation.
Springer, 1986, p. 445.
*/

package random

import "fmt"
import "math"

// GetTSampler returns a sampler function for the Student-t distribution.
//
// Distribution: T
//
// Parameters:
//   - nu: degrees of freedom (>0)
//
// The probability density function is:
//
//	f(x) = Γ((nu+1)/2) / (Γ(nu/2) * sqrt(pi * nu)) *
//	       (1 + x^2/nu)^(-(nu+1)/2),  for -∞ < x < ∞
//
// Example usage:
//
//	sampler, _ := r.GetTSampler(2.0)
//	x := sampler()
func (rng *Random) GetTSampler(nu float64) (func() float64, error) {
	if nu <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: nu must be positive")
	}
	return func() float64 {
		return rng.Normal() / math.Sqrt(rng.Gamma(nu/2.0)*2.0/nu)
	}, nil
}
