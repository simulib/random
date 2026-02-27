// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Lognormal random variate generator.

This file implements sampling from the lognormal distribution by
exponentiating a normally distributed random variate. If Z is normally
distributed with mean mu and standard deviation sigma, then exp(Z)
follows a lognormal distribution with the same parameters.

Reference:

C. Forbes, M. Evans, N. Hastings, B. Peacock.
Statistical Distributions.
Wiley, 2011 (4th ed.), Chapter 29.
*/

package random

import "fmt"
import "math"

// GetLognormalSampler returns a sampler function for the lognormal distribution.
//
// Distribution: Lognormal
//
// Parameters:
//   - mu: mean of the logarithm of the random variable
//   - sigmaL: standard deviation of the logarithm of the random variable (>0)
//
// The probability density function is:
//
//	f(x) = 1 / (sigma x sqrt(2*pi)) *
//	       exp(-(ln(x)-mu)^2 / (2*sigma^2)),  for x > 0
//
// Example:
//
//	sampler, _ := r.GetLognormalSampler(1.0, 2.0)
//	x := sampler()
func (rng *Random) GetLognormalSampler(mu, sigma float64) (func() float64, error) {
	if sigma <= 0.0 { // this is not strictly necessary, but most definitions assume sigma > 0
		return nil, fmt.Errorf("invalid parameter: sigma must be positive")
	}
	return func() float64 {
		return math.Exp(mu + sigma*rng.Normal())
	}, nil
}
