// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Poisson random variate generator.

This file implements sampling from the Poisson distribution using two methods:

1) The inversion (multiplicative product-of-uniforms) method for small mean values
   (λ ≤ 30).
2) The PTRS (Poisson Transformed Rejection with Squeeze) method for larger λ

Reference:

L. Devroye.
Non-Uniform Random Variate Generation.
Springer, 1986, p. 504.

G. Hörmann.
The Transformed Rejection Method for Generating Poisson Random Variables.
Insurance: Mathematics and Economics 12(1):39–45, 1993.
*/

package random

import (
	"fmt"
	"math"
)

// GetPoissonSampler returns a sampler function for the Poisson distribution.
//
// Distribution: Poisson
//
// Parameters:
//   - lambda: mean (≥0)
//
// The probability mass function is:
//
//	f(x) = λ^x * exp(-λ) / x!,  for x = 0, 1, 2, ...
//	     = 0,                   otherwise
//
// Example:
//
//	sampler, _ := r.GetPoissonSampler(3.0)
//	x := sampler()
func (rng *Random) GetPoissonSampler(lambda float64) (func() int64, error) {
	if lambda < 0.0 {
		return nil, fmt.Errorf("invalid parameter: lambda must be non-negative")
	}
	if lambda < 30 { // This is Knuth's algorithm
		P := newPoissonProduct(lambda)
		return func() int64 {
			return P.sample(rng)
		}, nil
	}
	P := newPoissonPTRS(lambda)
	return func() int64 {
		return P.sample(rng)
	}, nil
}

// Poisson returns a Poisson distributed random variable with rate lambda
func (rng *Random) Poisson(lambda float64) int64 {
	if lambda < 30 {
		P := newPoissonProduct(lambda)
		return P.sample(rng)
	}
	P := newPoissonPTRS(lambda)
	return P.sample(rng)
}

type poissonProduct struct {
	expNegLambda float64
}

func newPoissonProduct(lambda float64) *poissonProduct {
	P := &poissonProduct{expNegLambda: math.Exp(-lambda)}
	return P
}

func (P *poissonProduct) sample(rng *Random) int64 {
	var k int64
	prod := 1.0
	for {
		prod *= rng.Uniform()
		if prod <= P.expNegLambda {
			return k
		}
		k++
	}
}

type PoissonPTRS struct {
	lambda, slam, loglam, a, b, invAlpha, vr float64
}

func newPoissonPTRS(lambda float64) *PoissonPTRS {
	P := &PoissonPTRS{}
	P.lambda = lambda
	P.slam = math.Sqrt(lambda)
	P.loglam = math.Log(lambda)
	P.b = 0.931 + 2.53*P.slam
	P.a = -0.059 + 0.02483*P.b
	P.invAlpha = 1.1239 + 1.1328/(P.b-3.4)
	P.vr = 0.9277 - 3.6224/(P.b-2.0)
	return P
}

func (p *PoissonPTRS) sample(rng *Random) int64 {
	for {
		u := rng.Uniform() - 0.5
		v := 1.0 - rng.Uniform()
		us := 0.5 - math.Abs(u)
		if us == 0.0 {
			continue
		}
		k := int64(math.Floor((2.0*p.a/us+p.b)*u + p.lambda + 0.43))
		if us >= 0.07 && v <= p.vr {
			return k
		}
		if k < 0 || (us < 0.013 && v > us) {
			continue
		}
		lhs := math.Log(v) + math.Log(p.invAlpha) - math.Log(p.a/(us*us)+p.b)
		rhs := -p.lambda + float64(k)*p.loglam - logFactorial(k)
		if lhs <= rhs {
			return k
		}
	}
}

var logFactTable = []float64{
	0.0,
	0.0,
	0.6931471805599453,
	1.791759469228055,
	3.1780538303479458,
	4.787491742782046,
	6.579251212010101,
	8.525161361065415,
	10.60460290274525,
	12.801827480081469,
	15.104412573075516,
	17.502307845873887,
	19.987214495661885,
	22.552163853123425,
	25.19122118273868,
	27.89927138384089,
	30.671860106080672,
	33.50507345013689,
	36.39544520803305,
	39.339884187199495,
	42.335616460753485,
	45.38013889847691,
	48.47118135183523,
	51.60667556776438,
	54.78472939811232,
	58.00360522298052,
	61.261701761002,
	64.55753862700634,
	67.88974313718154,
	71.25703896716801,
	74.65823634883016,
	78.0922235533153,
	81.55795945611504,
	85.05446701758152,
	88.58082754219768,
	92.1361756036871,
	95.7196945421432,
	99.33061245478743,
	102.96819861451381,
	106.63176026064346,
	110.32063971475739,
	114.0342117814617,
	117.77188139974507,
	121.53308151543864,
	125.3172711493569,
	129.12393363912722,
	132.95257503561632,
	136.80272263732635,
	140.67392364823425,
	144.5657439463449,
	148.47776695177302,
	152.40959258449735,
	156.3608363030788,
	160.3311282166309,
	164.32011226319517,
	168.32744544842765,
	172.3527971391628,
	176.39584840699735,
	180.45629141754378,
	184.53382886144948,
	188.6281734236716,
	192.7390472878449,
	196.86618167289,
	201.00931639928152}

// logFactorial uses exact table values of log(n!) for small n, and Stirling's
// approximation for higher values.
func logFactorial(n int64) float64 {
	const logpi = 0.9189385332046727
	if n < int64(len(logFactTable)) {
		return logFactTable[n]
	}
	x := float64(n)
	return logpi + (x+0.5)*math.Log(x) - x + (1.0/12.0-1/(360.0*x*x))/x
}
