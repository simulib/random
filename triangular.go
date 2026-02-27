// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Triangular random variate generator.

This file implements sampling from the triangular distribution
using inversion of the cumulative distribution function.

*/

package random

import "fmt"
import "math"

// GetTriangularSampler returns a sampler function for the Triangular distribution.
//
// Distribution: Triangular
//
// Parameters:
//   - a: minimum
//   - b: maximum ( b > a)
//   - c: mode (a ≤ c ≤ b)
//
// The density function is
//
//	f(x) = 2 (x - a) / ((b - a)(c - a)),   for a ≤ x ≤ c
//	     = 2 (b - x) / ((b - a)(b - c)),   for c < x ≤ b
//	     = 0,                              otherwise
//
// Example usage:
//
//	sampler, _ := r.GetTriangularSampler(1.0, 3.0, 2.0)
//	x := sampler()
func (rng *Random) GetTriangularSampler(a, b, c float64) (func() float64, error) {
	if a >= b {
		return nil, fmt.Errorf("invalid parameters: a must be strictly smaller than b")
	}
	if !(a <= c && c <= b) {
		return nil, fmt.Errorf("invalid parameters: a must be smaller than c, c must be smaller than b")
	}
	fc := (c - a) / (b - a)
	span := b - a
	return func() float64 {
		u := rng.Uniform()
		if u < fc {
			return a + math.Sqrt(u*span*(c-a))
		}
		return b - math.Sqrt((1-u)*span*(b-c))
	}, nil
}
