// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Pareto random variate generator.

This file implements sampling from the Pareto distribution using
the inversion method: if U ~ Uniform(0,1), then
X = a / (1 - U)^(1/c) follows a Pareto(a, c) distribution.

*/

package random

import "fmt"
import "math"

// GetParetoSampler returns a sampler function for the Pareto distribution.
//
// Distribution: Pareto
//
// Parameters:
//   - a: scale parameter a (>0)
//   - c: shape parameter c (> 0).
//
// The probability density function is:
//
//	f(x) = c * a^c / x^(c + 1),  for x >= a
//	     = 0,                     otherwise
//
// Example:
//
//	sampler, _ := r.GetParetoSampler(1.0, 2.0)
//	x := sampler()
func (rng *Random) GetParetoSampler(a float64, c float64) (func() float64, error) {
	if a <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: a must be positive")
	}
	if c <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: c must be positive")
	}
	return func() float64 {
		return a * math.Pow(1.0-rng.Uniform(), -1.0/c)
	}, nil
}
