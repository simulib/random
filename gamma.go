// Copyright (c) 2025 Dieter Fiems
// SPDX-License-Identifier: MIT

package random

/*
Gamma random variate generator.

This file implements sampling from the Gamma distribution using the
acceptance–rejection method of Marsaglia and Tsang.

Reference:

G. Marsaglia and W. W. Tsang.
A Simple Method for Generating Gamma Variables.
ACM Transactions on Mathematical Software 26(3):363–372, 2000.
*/

import (
	"fmt"
	"math"
)

// GetGammaSampler returns a sampler function for the Gamma distribution.
//
// Distribution: Gamma
//
// Parameters:
//   - alpha: shape parameter (>0)
//   - theta: scale parameter (>0)
//
// The probability density function is:
//
//	f(x) = x^(alpha-1) * exp(-x/theta) / (Gamma(alpha) * theta^alpha),  for x ≥ 0
//	     = 0,                                                          otherwise
//
// where Gamma() is the Gamma function.
//
// Example:
//
//	sampler, _ := r.GetGammaSampler(1.0, 2.0)
//	x := sampler()
func (rng *Random) GetGammaSampler(alpha float64, theta float64) (func() float64, error) {
	if theta < 0.0 {
		return nil, fmt.Errorf("invalid parameter: theta must be non-negative")
	}
	if alpha <= 0.0 {
		return nil, fmt.Errorf("invalid parameter: alpha must be positive")
	}
	return func() float64 { return theta * rng.Gamma(alpha) }, nil
}

// Gamma returns a gamma random variate with shape parameter alpha > 0 and scale parameter 1
func (rng *Random) Gamma(alpha float64) float64 {
	if alpha < 1.0 {
		return rng.gamma1(alpha+1) * math.Pow(rng.Uniform(), 1.0/alpha)
	}
	if alpha == 1.0 {
		return rng.Exponential()
	}
	return rng.gamma1(alpha)
}

// gamma1 generates Gamma(alpha,1) for alpha >= 1.
func (rng *Random) gamma1(alpha float64) float64 {
	d := alpha - 1.0/3.0
	c := 1.0 / math.Sqrt(9.0*d)
	var x, u, v float64
	for {
		x = rng.Normal()
		v = 1.0 + c*x
		if v <= 0 {
			continue
		}
		v = v * v * v
		u = rng.Uniform()
		if (u < 1.0-0.0331*(x*x)*(x*x)) || (math.Log(u) < 0.5*x*x+d*(1.0-v+math.Log(v))) {
			// note that the second condition is not evaluated if the first holds
			return d * v
		}
	}
}
