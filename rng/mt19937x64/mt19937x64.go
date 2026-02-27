// Copyright (c) 2025 Dieter Fiems
// SPDX-License-Identifier: MIT

// This is the 64-bit Mersenne Twister from Nishimura (2000), which is
// a 64-bit version of the Mersenne Twister of Matsumoto and Nishimura (1998)
//
//	T. Nishimura, ``Tables of 64-bit Mersenne Twisters''
//	ACM Transactions on Modeling and Computer Simulation 10. (2000) 348--357.
//
//	M. Matsumoto and T. Nishimura, ``Mersenne Twister: a 623-dimensionally
//	equidistributed uniform pseudorandom number generator''
//	ACM Transactions on Modeling and Computer Simulation 8. (Jan. 1998) 3--30.
//
// This generator uses the parameter set from the 2004 version of the authors' code
package mt19937x64

const (
	mm int    = 156
	nn int    = 312
	u  int    = 29
	s  int    = 17
	t  int    = 37
	l  int    = 43
	a  uint64 = 0xB5026F5AA96619E9
	b  uint64 = 0x71D67FFFEDA60000
	c  uint64 = 0xFFF7EEE000000000
	d  uint64 = 0x5555555555555555
	mU uint64 = 0xFFFFFFFF80000000 /* most significant 64-r bits */
	mL uint64 = 0x7FFFFFFF         /* least significant r bits */
)

// This is the 64-bit Mersenne Twister from Nishimura (2000), which is
// a 64-bit version of the Mersenne Twister of Matsumoto and Nishimura (1998)
//
//	T. Nishimura, ``Tables of 64-bit Mersenne Twisters''
//	ACM Transactions on Modeling and Computer Simulation 10. (2000) 348--357.
//
//	M. Matsumoto and T. Nishimura, ``Mersenne Twister: a 623-dimensionally
//	equidistributed uniform pseudorandom number generator''
//	ACM Transactions on Modeling and Computer Simulation 8. (Jan. 1998) 3--30.
//
// This generator uses the parameter set from the 2004 version of the authors' code
type Mt19937x64 struct {
	w     [nn]uint64
	i     int
	cycle int
}

// New creates a new Mersenne Twister random number generator with the given seed
func New(seed uint64) *Mt19937x64 {
	r := new(Mt19937x64)
	r.Seed(seed)
	return r
}

// Seed reinitialises the Mersenne Twister random number generator with the given seed
func (r *Mt19937x64) Seed(seed uint64) {
	var i int
	r.w[0] = seed
	for i = 1; i < nn; i++ {
		r.w[i] = 6364136223846793005*(r.w[i-1]^(r.w[i-1]>>62)) + uint64(i)
	}
	r.cycle = 1
	r.i = 0
}

// Uint64 generates a random number in the  [0, 2^64-1]-interval
func (r *Mt19937x64) Uint64() uint64 {
	mag01 := [2]uint64{0, a}
	var x uint64
	switch r.cycle {
	case 1:
		x = (r.w[r.i] & mU) | (r.w[r.i+1] & mL)
		r.w[r.i] = r.w[r.i+mm] ^ (x >> 1) ^ mag01[int(x&1)]
		x = r.w[r.i]
		r.i++
		if r.i == nn-mm {
			r.cycle = 2
		}
		break
	case 2:
		x = (r.w[r.i] & mU) | (r.w[r.i+1] & mL)
		r.w[r.i] = r.w[r.i+mm-nn] ^ (x >> 1) ^ mag01[int(x&1)]
		x = r.w[r.i]
		r.i++
		if r.i == nn-1 {
			r.cycle = 3
		}
		break
	case 3:
		x = (r.w[nn-1] & mU) | (r.w[0] & mL)
		r.w[nn-1] = r.w[mm-1] ^ (x >> 1) ^ mag01[int(x&1)]
		x = r.w[nn-1]
		r.cycle = 1
		r.i = 0
		break
	}
	x ^= (x >> u) & d
	x ^= (x << s) & b
	x ^= (x << t) & c
	x ^= x >> l
	return x
}
