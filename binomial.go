// Copyright (c) 2026 Dieter Fiems
// SPDX-License-Identifier: MIT

/*
Binomial random variate generator.

This file implements algorithms for generating binomially distributed
random variates. The choice of algorithm depends on the values of n and p.

For n·p ≥ 20, the BTPE method of Kachitvichyanukul and Schmeiser is used.
For n·p < 20 and n < 1000, the inversion method for discrete random
variables is used.
For n·p < 20 and n ≥ 1000, the second waiting time method of Devroye
is used.

References:

[1] V. Kachitvichyanukul, B. W. Schmeiser.
    Binomial random variate generation.
    Communications of the ACM, 31(2):216–222, 1988.

[2] L. Devroye.
    Non-Uniform Random Variate Generation.
    Springer, 1986, p. 525.
*/

package random

import (
	"fmt"
	"math"
)

// GetBinomialSampler returns a sampler function for the binomial distribution.
//
// Distribution: Binomial
//
// Parameters:
//   - n: number of trials (>0, integer)
//   - p: success probability (0 < p < 1).
//
// The probability mass function is
//
//	f(x) = C(n, x) * p^x * (1-p)^(n-x),  for x = 0, ..., n
//		 = 0,                        otherwise
//
// where C(n, x) = n! / (x! (n-x)!) is the binomial coefficient.
//
// Example:
//
//	sampler, _ := r.GetBinomialSampler(18, 0.3)
//	x := sampler()
func (rng *Random) GetBinomialSampler(n int64, p float64) (func() int64, error) {
	if n <= 0 {
		return nil, fmt.Errorf("invalid parameter: n must be positive")
	}
	if p <= 0 || p >= 1.0 {
		return nil, fmt.Errorf("invalid parameter: p must be in (0,1)")
	}
	if float64(n)*min(p, 1-p) >= 20 {
		B := newBinomialBTPE(n, p)
		return func() int64 {
			return B.sample(rng)
		}, nil
	}
	if n < 1000 {
		B := newBinomialInversion(n, p)
		return func() int64 {
			return B.sample(rng)
		}, nil
	}
	B := newBinomialWaitingTime(n, p)
	return func() int64 {
		return B.sample(rng)
	}, nil
}

type binomialBTPE struct {
	r, p, q, nrq, fM, p1, p2, p3, p4, xM, xL, xR, c, lambdaL, lambdaR float64
	M, n                                                              int64
}

// newBinomialBTPE initialises the BTPE parameters
func newBinomialBTPE(n int64, p float64) *binomialBTPE {
	// step 0
	bin := &binomialBTPE{}
	bin.p = p
	bin.n = n
	bin.r = min(p, 1-p)
	bin.q = 1 - bin.r
	bin.nrq = float64(bin.n) * bin.r * bin.q
	bin.fM = float64(n)*bin.r + bin.r
	bin.M = int64(bin.fM)
	bin.p1 = math.Floor(2.195*math.Sqrt(bin.nrq)-4.6*bin.q) + 0.5
	bin.xM = float64(bin.M) + 0.5
	bin.xL = bin.xM - bin.p1
	bin.xR = bin.xM + bin.p1
	bin.c = 0.134 + 20.5/(15.3+float64(bin.M))
	a := (bin.fM - bin.xL) / (bin.fM - bin.xL*bin.r)
	bin.lambdaL = a * (1 + a/2)
	a = (bin.xR - bin.fM) / (bin.xR * bin.q)
	bin.lambdaR = a * (1 + a/2)
	bin.p2 = bin.p1 * (1 + 2.0*bin.c)
	bin.p3 = bin.p2 + bin.c/bin.lambdaL
	bin.p4 = bin.p3 + bin.c/bin.lambdaR
	return bin
}

// sample samples using BTPE
func (bin binomialBTPE) sample(rng *Random) (y int64) {
	var u, v, x, k, s, a, F, rho, t, A, x1, x2, f1, f2, z, z2, w, w2 float64
	var i int64
step1:
	u = bin.p4 * rng.Uniform()
	v = rng.Uniform()
	if u > bin.p1 {
		goto step2
	}
	y = int64(math.Floor(bin.xM - bin.p1*v + u))
	goto step6
step2:
	if u > bin.p2 {
		goto step3
	}
	x = bin.xL + (u-bin.p1)/bin.c
	v = v*bin.c + 1.0 - math.Abs(bin.xM-x)/bin.p1
	if v > 1.0 {
		goto step1
	}
	y = int64(math.Floor(x))
	goto step5
step3:
	if u > bin.p3 {
		goto step4
	}
	if v == 0.0 {
		goto step1
	}
	y = int64(math.Floor(bin.xL + math.Log(v)/bin.lambdaL))
	if y < 0 {
		goto step1
	}
	v *= (u - bin.p2) * bin.lambdaL
	goto step5
step4:
	y = int64(math.Floor(bin.xR - math.Log(v)/bin.lambdaR))
	if y > bin.n {
		goto step1
	}
	v = v * (u - bin.p3) * bin.lambdaR
step5:
	k = math.Abs(float64(y - bin.M))
	if k > 20.0 && k < bin.nrq/2.0-1.0 {
		goto step52
	}
	//step51:
	s = bin.r / bin.q
	a = s * float64(bin.n+1)
	F = 1.0
	if bin.M < y {
		for i = bin.M + 1; i <= y; i++ {
			F *= a/float64(i) - s
		}
	}
	if bin.M > y {
		for i = y + 1; i <= bin.M; i++ {
			F /= a/float64(i) - s
		}
	}
	if v > F {
		goto step1
	}
	goto step6
step52:
	rho = (k / bin.nrq) * ((k*(k/3.0+0.625)+0.16666666666666666)/bin.nrq + 0.5)
	t = -k * k / (2.0 * bin.nrq)
	if v == 0.0 {
		goto step1
	}
	A = math.Log(v)
	if A < t-rho {
		goto step6
	}
	if A > t+rho {
		goto step1
	}
	//step53:
	x1 = float64(y) + 1.0
	f1 = float64(bin.M + 1)
	z = float64(bin.n + 1 - bin.M)
	w = float64(bin.n - y + 1)
	x2 = x1 * x1
	f2 = f1 * f1
	z2 = z * z
	w2 = w * w
	if A > (bin.xM*math.Log(f1/x1) + (float64(bin.n-bin.M)+0.5)*math.Log(z/w) +
		float64(y-bin.M)*math.Log(w*bin.r/(x1*bin.q)) +
		(13680.-(462.-(132.-(99.-140./f2)/f2)/f2)/f2)/f1/166320. +
		(13680.-(462.-(132.-(99.-140./z2)/z2)/z2)/z2)/z/166320. +
		(13680.-(462.-(132.-(99.-140./x2)/x2)/x2)/x2)/x1/166320. +
		(13680.-(462.-(132.-(99.-140./w2)/w2)/w2)/w2)/w/166320.) {
		goto step1
	}
step6:
	if bin.p > 0.5 {
		y = bin.n - y
	}
	return
}

type binomialInversion struct {
	r, p, q, p0 float64
	n           int64
}

// newBinomialInversion initialises the parameters for the inversion method
func newBinomialInversion(n int64, p float64) *binomialInversion {
	bin := &binomialInversion{}
	bin.p = p
	bin.n = n
	bin.r = min(p, 1.0-p)
	bin.q = 1.0 - bin.r
	bin.p0 = math.Exp(float64(n) * math.Log(bin.q))
	return bin
}

// sample samples using the inversion method
func (bin binomialInversion) sample(rng *Random) (y int64) {
	y = 0
	P := bin.p0
	U := rng.Uniform()
	for U > P && y < bin.n {
		y++
		U -= P
		P = (float64(bin.n-y+1) * bin.r * P) / (float64(y) * bin.q)
	}
	if bin.p > 0.5 {
		y = bin.n - y
	}
	return
}

type binomialWaitingTime struct {
	p, q float64
	n    int64
}

// newBinomialWaitingTime initialises the parameters for the waiting time method
func newBinomialWaitingTime(n int64, p float64) *binomialWaitingTime {
	bin := &binomialWaitingTime{}
	bin.p = p
	bin.n = n
	bin.q = -math.Log(max(p, 1.0-p))
	return bin
}

// sample samples using the waiting time method
func (bin binomialWaitingTime) sample(rng *Random) (y int64) {
	y = 0
	sum := 0.0
	for {
		sum += rng.Exponential() / float64(bin.n-y)
		y++
		if sum > bin.q {
			y = y - 1
			if bin.p > 0.5 {
				y = bin.n - y
			}
			return
		}
	}
}
