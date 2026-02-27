// Copyright (c) 2025 Dieter Fiems
// SPDX-License-Identifier: MIT

// Package pcg64dxsm implements a PCG (Permuted Congruential Generator) random
// number generator as described in:
//
//	Melissa E. O’Neill,
//	"PCG: A Family of Simple Fast Space-Efficient Statistically Good
//	Algorithms for Random Number Generation",
//	HMC-CS-2014-0905, Harvey Mudd College, Claremont, CA, 2014.
//
// This package provides a 128-bit LCG with DXSM output permutation,
// referred to as "pcg_engines::setseq_dxsm_128_64" in the reference implementation.
package pcg64dxsm

import "math/bits"

// Pcg64dxsm implements the PCG 128/64 DXSM generator described in
//
//	Melissa E. O’Neill,
//	"PCG: A Family of Simple Fast Space-Efficient Statistically Good
//	Algorithms for Random Number Generation",
//	HMC-CS-2014-0905, Harvey Mudd College, Claremont, CA, 2014.
//
// It uses a 128-bit linear congruential state and the DXSM output
// permutation. The generator supports jump-ahead for stream splitting.
type Pcg64dxsm struct {
	hi, lo       uint64
	incHi, incLo uint64
}

// New creates a new pcg64dxsm random number generator with the given (seed1,seed2)
func New(seed1, seed2 uint64) *Pcg64dxsm {
	r := new(Pcg64dxsm)
	r.Seed(seed1, seed2)
	return r
}

// Seed reinitialises the pcg64dxsm random number generator with the given (seed1,seed2)
func (r *Pcg64dxsm) Seed(seed1, seed2 uint64) {
	r.hi = 0
	r.lo = 0
	_ = r.Uint64()
	var c uint64
	r.lo, c = bits.Add64(r.lo, seed2, 0)
	r.hi, _ = bits.Add64(r.hi, seed1, c)
	_ = r.Uint64()
}

// Uint64 generates a random number in the  [0, 2^64-1]-interval
func (r *Pcg64dxsm) Uint64() uint64 {
	const (
		mulHi    = 2549297995355413924
		mulLo    = 4865540595714422341
		incHi    = 6364136223846793005
		incLo    = 1442695040888963407
		cheapMul = 0xda942042e4dd58b5
	)
	// update state <- state * mul + inc
	hi, lo := bits.Mul64(r.lo, mulLo)
	hi += r.hi*mulLo + r.lo*mulHi
	var c uint64
	lo, c = bits.Add64(lo, incLo, 0)
	hi, _ = bits.Add64(hi, incHi, c)
	r.lo, r.hi = lo, hi

	// dxsm output function
	hi ^= hi >> 32
	hi *= cheapMul
	hi ^= hi >> 48
	hi *= lo | 1
	return hi
}

// Jump is equivalent to 2^64 calls of Uint64()
func (r *Pcg64dxsm) Jump() {
	const (
		mulHi = 6848745282420445116
		mulLo = 1
		incHi = 4350103176503831857
		incLo = 0
	)
	hi, lo := bits.Mul64(r.lo, mulLo)
	hi += r.hi*mulLo + r.lo*mulHi
	var c uint64
	r.lo, c = bits.Add64(lo, incLo, 0)
	r.hi, _ = bits.Add64(hi, incHi, c)
}

// LongJump is equivalent to 2^96 calls of Uint64()
func (r *Pcg64dxsm) Longjump() {
	const (
		mulHi = 6038640712990326784
		mulLo = 1
		incHi = 10074320579927736320
		incLo = 0
	)
	hi, lo := bits.Mul64(r.lo, mulLo)
	hi += r.hi*mulLo + r.lo*mulHi
	var c uint64
	r.lo, c = bits.Add64(lo, incLo, 0)
	r.hi, _ = bits.Add64(hi, incHi, c)
}
