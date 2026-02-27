// Copyright (c) 2025 Dieter Fiems
// SPDX-License-Identifier: MIT

package xoshiro256starstar

import "github.com/simulib/random/rng/splitmix64"

/*
 * This is the Xoshiro256** random number generator described in
 *
 * David Blackman, Sebastiano Vigna.
 * Scrambled Linear Pseudorandom Number Generators.
 * ACM Trans. Math. Softw., 2021.
 *
 * Initialization uses the splitmix generator.
 *
 */

type Xoshiro256starstar struct {
	s [4]uint64
}

// New creates a new xoshiro256++ random number generator with the given seed
func New(seed uint64) *Xoshiro256starstar {
	m := new(Xoshiro256starstar)
	m.Seed(seed)
	return m
}

// Seed reinitialises the xoshiro256++ random number generator with the given seed
func (r *Xoshiro256starstar) Seed(seed uint64) {
	rng := splitmix64.New(seed)
	r.s[0] = rng.Uint64()
	r.s[1] = rng.Uint64()
	r.s[2] = rng.Uint64()
	r.s[3] = rng.Uint64()
}

// Uint64 generates a random number in the  [0, 2^64-1]-interval
func (r *Xoshiro256starstar) Uint64() uint64 {
	x := 5 * r.s[1]
	result := ((x << 7) | (x >> 57)) * 9
	t := r.s[1] << 17
	r.s[2] ^= r.s[0]
	r.s[3] ^= r.s[1]
	r.s[1] ^= r.s[2]
	r.s[0] ^= r.s[3]
	r.s[2] ^= t
	r.s[3] = (r.s[3] << 45) | (r.s[3] >> 19)
	return result
}

// Jump is equivalent to 2^128 calls to Uint64()
func (r *Xoshiro256starstar) Jump() {
	jump := [...]uint64{0x180ec6d33cfd0aba, 0xd5a61266f0c9392c, 0xa9582618e03fc9aa, 0x39abdc4529b1661c}
	var s0, s1, s2, s3 uint64
	for i := 0; i < 4; i++ {
		for b := 0; b < 64; b++ {
			if jump[i]&(uint64(1)<<b) != 0 {
				s0 ^= r.s[0]
				s1 ^= r.s[1]
				s2 ^= r.s[2]
				s3 ^= r.s[3]
			}
			_ = r.Uint64()
		}
	}
	r.s[0] = s0
	r.s[1] = s1
	r.s[2] = s2
	r.s[3] = s3
}

// Longjump is equivalent to 2^192 calls to Uint64()
func (r *Xoshiro256starstar) Longjump() {
	jump := [...]uint64{0x76e15d3efefdcbbf, 0xc5004e441c522fb3, 0x77710069854ee241, 0x39109bb02acbe635}
	var s0, s1, s2, s3 uint64
	for i := 0; i < 4; i++ {
		for b := 0; b < 64; b++ {
			if jump[i]&(uint64(1)<<b) != 0 {
				s0 ^= r.s[0]
				s1 ^= r.s[1]
				s2 ^= r.s[2]
				s3 ^= r.s[3]
			}
			_ = r.Uint64()
		}
	}
	r.s[0] = s0
	r.s[1] = s1
	r.s[2] = s2
	r.s[3] = s3
}
