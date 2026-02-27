// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Geometric random variate generator.

This file implements sampling from the geometric distribution as
described by Devroye. The implementation uses the ceil of an
exponential random variate.

Reference:

L. Devroye.
Non-Uniform Random Variate Generation.
Springer, 1986, p. 500.
*/

package random

import (
	"fmt"
	"math"
)

// GetGeometricSampler returns a sampler function for the geometric distribution.
//
// Distribution: Geometric
//
// Parameters:
//   - p: success probability (0 < p ≤ 1).
//
// The probability mass function is:
//
//	p(x) = p * (1-p)^(x-1),  for x = 1, 2, 3, ...
//	     = 0,                  otherwise
//
// Example:
//
//	sampler, _ := r.GetGeometricSampler(0.3)
//	x := sampler()
func (rng *Random) GetGeometricSampler(p float64) (func() int64, error) {
	if p <= 0 || p > 1 {
		return nil, fmt.Errorf("invalid parameter: p must be in the interval (0,1]")
	}
	if p == 1.0 {
		return func() int64 { return 1 }, nil
	}
	scale := -1.0 / math.Log(1.0-p)
	return func() int64 {
		val := rng.Exponential() * scale
		if val >= math.MaxInt64 {
			return math.MaxInt64
		}
		return int64(math.Ceil(val))
	}, nil
}
