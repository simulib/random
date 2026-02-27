// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Logarithmic series random variate generator.

This file implements sampling from the logarithmic series distribution
using Kemp’s second accelerated generator.

References:

L. Devroye.
Non-Uniform Random Variate Generation.
Springer, 1986, p. 548.

A. W. Kemp.
Efficient Generation of Logarithmically Distributed Random Variables.
Journal of the Royal Statistical Society, Series C (Applied Statistics),
30(3): 249–253, 1981.
*/

package random

import (
	"fmt"
	"math"
)

// GetLogarithmicSampler returns a sampler function for the logarithmic distribution.
//
// Distribution: Logarithmic series distribution
//
// Parameters:
//   - c: shape parameter (0<c<1)
//
// The probability mass function is:
//
//	f(x) = c^x / (x * log(1/(1-c))),  for x = 1, 2, 3, ...
//	     = 0,                         otherwise
//
// Example:
//
//	sampler, _ := r.GetLogarithmicSampler(0.7)
//	x := sampler()
func (rng *Random) GetLogarithmicSampler(c float64) (func() int64, error) {
	if c <= 0.0 || c >= 1.0 {
		return nil, fmt.Errorf("invalid parameter: c must be in the interval (0,1)")
	}
	r := math.Log(1.0 - c)
	return func() int64 {
		v := rng.Uniform()
		if v >= c {
			return 1
		}
		q := 1.0 - math.Exp(r*rng.Uniform())
		if v <= q*q {
			return int64(math.Floor(1.0 + math.Log(v)/math.Log(q)))
		}
		if v <= q {
			return 2
		}
		return 1
	}, nil

}
