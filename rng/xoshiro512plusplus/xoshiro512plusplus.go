// Copyright (c) 2025 Dieter Fiems
// SPDX-License-Identifier: MIT

package xoshiro512plusplus

import "github.com/simulib/random/rng/splitmix64"

/*
 * This is the Xoshiro512++ random number generator described in
 *
 * David Blackman, Sebastiano Vigna.
 * Scrambled Linear Pseudorandom Number Generators.
 * ACM Trans. Math. Softw., 2021.
 *
 * Initialization uses the splitmix generator.
 *
 */
type Xoshiro512plusplus struct {
	s [8]uint64
}

// New creates a new xoshiro256++ random number generator with the given seed
func New(seed uint64) *Xoshiro512plusplus {
	m := new(Xoshiro512plusplus)
	m.Seed(seed)
	return m
}

// Seed reinitialises the xoshiro256++ random number generator with the given seed
func (r *Xoshiro512plusplus) Seed(seed uint64) {
	rng := splitmix64.New(seed)
	for i := 0; i < 8; i++ {
		r.s[i] = rng.Uint64()
	}
}

// Uint64 generates a random number in the  [0, 2^64-1]-interval
func (r *Xoshiro512plusplus) Uint64() uint64 {
	x := r.s[0] + r.s[2]
	result := ((x << 17) | (x >> 47)) + r.s[2]
	t := r.s[1] << 11
	r.s[2] ^= r.s[0]
	r.s[5] ^= r.s[1]
	r.s[1] ^= r.s[2]
	r.s[7] ^= r.s[3]
	r.s[3] ^= r.s[4]
	r.s[4] ^= r.s[5]
	r.s[0] ^= r.s[6]
	r.s[6] ^= r.s[7]
	r.s[6] ^= t
	r.s[7] = (r.s[7] << 21) | (r.s[7] >> 43)
	return result
}

// Jump is equivalent to 2^256 calls to Uint64()
func (r *Xoshiro512plusplus) Jump() {
	jump := [8]uint64{0x33ed89b6e7a353f9, 0x760083d7955323be, 0x2837f2fbb5f22fae, 0x4b8c5674d309511c,
		0xb11ac47a7ba28c25, 0xf1be7667092bcc1c, 0x53851efdb6df0aaf, 0x1ebbc8b23eaf25db}
	var t [8]uint64
	for i := 0; i < 8; i++ {
		for b := 0; b < 64; b++ {
			if (jump[i] & (uint64(1) << b)) != 0 {
				for w := 0; w < 8; w++ {
					t[w] ^= r.s[w]
				}
			}
			_ = r.Uint64()
		}
	}
	for i := 0; i < 8; i++ {
		r.s[i] = t[i]
	}
}

// Longjump is equivalent to 2^384 calls to Uint64()
func (r *Xoshiro512plusplus) Longjump() {
	jump := [8]uint64{0x11467fef8f921d28, 0xa2a819f2e79c8ea8, 0xa8299fc284b3959a, 0xb4d347340ca63ee1,
		0x1cb0940bedbff6ce, 0xd956c5c4fa1f8e17, 0x915e38fd4eda93bc, 0x5b3ccdfa5d7daca5}
	var t [8]uint64
	for i := 0; i < 8; i++ {
		for b := 0; b < 64; b++ {
			if (jump[i] & (uint64(1) << b)) != 0 {
				for w := 0; w < 8; w++ {
					t[w] ^= r.s[w]
				}
			}
			_ = r.Uint64()
		}
	}
	for i := 0; i < 8; i++ {
		r.s[i] = t[i]
	}
}
