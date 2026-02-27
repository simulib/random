package random

/*
Negative binomial and Pascal random variate generator.

This file implements sampling from the negative binomial distribution
using the Poisson–Gamma mixture representation: if
Y ~ Gamma(r, (1-p)/p) and X ~ Poisson(Y), then X follows
a NegativeBinomial(r, p) distribution.

For integer r, the negative binomial distribution is equivalent to the
Pascal distribution.

Reference:

L. Devroye.
Non-Uniform Random Variate Generation.
Springer, 1986, Example 1.5, p. 543.
*/

import "fmt"

// GetNegativeBinomialSampler returns a sampler function for the negative binomial distribution.
//
// Distribution: Negative Binomial
//
// Parameters:
//   - r: number of successes r (>0, can be real)
//   - p: success probability (0 < p ≤ 1)
//
// The probability mass function is:
//
//	f(x) = Gamma(r + x) / (Gamma(r) * x!) * p^r * (1-p)^x,  for x = 0, 1, 2, ...
//	     = 0,                                                otherwise
//
// Example:
//
//	sampler, _ := r.GetNegativeBinomial(5.0, 0.3)
//	x := sampler()
func (rng *Random) GetNegativeBinomialSampler(r, p float64) (func() int64, error) {
	if p <= 0 || p > 1 {
		return nil, fmt.Errorf("invalid parameter: p must be in the interval (0,1]")
	}
	if r <= 0 {
		return nil, fmt.Errorf("invalid parameter: r must be positive")
	}
	return func() int64 { // This is the Leger's algorithm (given in the answers in Knuth)
		X := rng.Gamma(r)
		return rng.Poisson(X * (1 - p) / p)
	}, nil
}

// GetPascalSampler returns a sampler function for the Pascal (negative binomial) distribution.
//
// Distribution: Pascal
//
// Parameters:
//   - r: number of successes r (>0, integer)
//   - p: success probability (0 < p ≤ 1)
//
// The probability mass function is:
//
//	f(x) = Gamma(r + x) / (Gamma(r) * Gamma(x + 1)) *
//	       p^r * (1 - p)^x,  for x = 0, 1, 2, ...
//	     = 0,                 otherwise
//
// Example:
//
//	sampler, _ := r.GetPascalSampler(50, 0.3)
//	x := sampler()
func (rng *Random) GetPascalSampler(r int64, p float64) (func() int64, error) {
	return rng.GetNegativeBinomialSampler(float64(r), p)
}
