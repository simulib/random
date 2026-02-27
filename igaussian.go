// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Inverse Gaussian (Wald) random variate generator.

This file implements sampling from the inverse Gaussian distribution
using the method of Michael, Schucany, and Haas.

Reference:

J. R. Michael, W. R. Schucany, R. W. Haas.
Generating Random Variates Using Transformations with Multiple Roots.
Communications in Statistics – Simulation and Computation, 5(2): 88–90, 1976.
*/

package random

import "fmt"
import "math"

// GetInverseGaussianSampler returns a sampler function for the inverse Gaussian distribution.
//
// Distribution: Inverse Gaussian
//
// Parameters:
//   - mu: mean (>0)
//   - lambda: shape parameter (>0)
//
// The probability density function is:
//
//	f(x) = sqrt(lambda / (2*pi*x^3)) * exp(-lambda*(x-mu)^2 / (2*mu^2*x)),  for x > 0
//	     = 0,                                                               otherwise
//
// Example:
//
//	sampler, _ := r.GetInverseGaussianSampler(1.0, 2.0)
//	x := sampler()
func (rng *Random) GetInverseGaussianSampler(mu float64, lambda float64) (func() float64, error) {
	if mu <= 0 {
		return nil, fmt.Errorf("invalid parameter: mu must be positive")
	}
	if lambda <= 0 {
		return nil, fmt.Errorf("invalid parameter: lambda must be positive")
	}
	return func() float64 {
		v := rng.Normal()
		w := mu * v * v
		c := mu / (2.0 * lambda)
		x := mu + c*(w-math.Sqrt(w*(4.0*lambda+w)))
		if rng.Uniform() < mu/(mu+x) {
			return x
		}
		return (mu * mu) / x
	}, nil
}
