// Copyright (c) 2025 Dieter Fiems
// SPDX-License-Identifier: MIT

package splitmix64

/*
 * This is the splitmix random number generator mix64variant13 described in
 *
 *	Guy L. Steele, Jr., Doug Lea, and Christine H. Flood.
 *	Fast Splittable Pseudorandom Number Generators.
 *	In: Proceedings of the 2014 ACM International Conference on Object Oriented Programming *
 *      Systems Languages & Applications (OOPSLA ’14). ACM, New York, NY, USA, 453–472, 2014.
 *
 */
type Splitmix struct {
	s uint64
}

// New creates a new splitmix random number generator with the given seed
func New(seed uint64) *Splitmix {
	m := new(Splitmix)
	m.Seed(seed)
	return m
}

// Seed reinitialises the splitmix random number generator with the given seed
func (r *Splitmix) Seed(seed uint64) {
	r.s = seed
}

// Uint64 generates a random number in the  [0, 2^64-1]-interval
func (r *Splitmix) Uint64() uint64 {
	r.s += 0x9e3779b97f4a7c15
	z := r.s
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}
