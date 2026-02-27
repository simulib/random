// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Logistic random variate generator.

This file implements sampling from the Logistic distribution using the
inversion method.

*/

package random

import "fmt"
import "math"

// GetLogisticSampler returns a sampler function for the logistic distribution.
//
// Distribution: Logistic
//
// Parameters:
//   - a : location parameter
//   - b : scale parameter  (>0)
//
// The probability density function is:
//
//	f(x) = exp(-(x-a)/b) / (b * (1 + exp(-(x-a)/b))^2),  for -∞ < x < ∞
//
// Example:
//
//	sampler, _ := r.GetLogisticSampler(1.0, 2.0)
//	x := sampler()
func (rng *Random) GetLogisticSampler(a float64, b float64) (func() float64, error) {
	if b <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: b must be positive")
	}
	return func() float64 {
		var x float64
		for {
			x = rng.Uniform()
			if (x > 0.0) && (x < 1.0) {
				return a + b*math.Log(x/(1-x))
			}
		}
	}, nil
}
