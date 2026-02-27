// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Uniform random variate generator.

This file implements sampling from the uniform distribution.
If U is a standard uniform number, then a + U*(b-a) is uniform
on the interval [a,b].

*/

package random

import "fmt"

// GetUniformSampler returns a sampler function for the Beta distribution.
//
// Distribution: Uniform (Rectangular)
//
// Parameters:
//   - a: minimum
//   - b: maximum (b>a)
//
// The density function is:
//
//	f(x) = 1/(b-a),   for a <= x <= b
//	     = 0,         otherwise
//
// Example usage:
//
//	sampler,_ := r.GetUniformSampler(1.0, 2.0)
//	x := sampler()
func (rng *Random) GetUniformSampler(a float64, b float64) (func() float64, error) {
	if a > b {
		return nil, fmt.Errorf("invalid parameter: a must be smaller than b")
	}
	return func() float64 { return a + (b-a)*rng.Uniform() }, nil
}

// Uniform returns a uniformly distributed random variate in [0,1)
func (rng *Random) Uniform() float64 {
	return float64(rng.Uint64()>>11) / 9007199254740992.0
}
