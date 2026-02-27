package philiox4x64

import "math/bits"

type Philox4x64 struct {
	counter [4]uint64
	key     [2]uint64
	buffer  [4]uint64
	idx     int
}

func New(key0, key1 uint64) *Philox4x64 {
	r := &Philox4x64{
		key: [2]uint64{key0, key1},
	}
	r.refill()
	return r
}

func (r *Philox4x64) Uint64() uint64 {
	if r.idx >= 4 {
		r.incrementCounter()
		r.refill()
	}
	x := r.buffer[r.idx]
	r.idx++
	return x
}

func round(c [4]uint64, k [2]uint64) [4]uint64 {
	const (
		M0 = 0xD2E7470EE14C6C93
		M1 = 0xCA5A826395121157
	)
	hi0, lo0 := bits.Mul64(M0, c[0])
	hi1, lo1 := bits.Mul64(M1, c[2])
	return [4]uint64{
		hi1 ^ c[1] ^ k[0],
		lo1,
		hi0 ^ c[3] ^ k[1],
		lo0,
	}
}

func bumpkey(k [2]uint64) [2]uint64 {
	const (
		W0 = 0x9E3779B97F4A7C15
		W1 = 0xBB67AE8584CAA73B
	)
	k[0] += W0
	k[1] += W1
	return k
}

func (r *Philox4x64) refill() {
	c := r.counter
	k := r.key
	c = round(c, k) // 1
	k = bumpkey(k)
	c = round(c, k) // 2
	k = bumpkey(k)
	c = round(c, k) // 3
	k = bumpkey(k)
	c = round(c, k) // 4
	k = bumpkey(k)
	c = round(c, k) // 5
	k = bumpkey(k)
	c = round(c, k) // 6
	k = bumpkey(k)
	c = round(c, k) // 7
	k = bumpkey(k)
	c = round(c, k) // 8
	k = bumpkey(k)
	c = round(c, k) // 9
	k = bumpkey(k)
	c = round(c, k) // 10

	r.buffer = c
	r.idx = 0
}

func (r *Philox4x64) incrementCounter() {
	r.counter[0]++
	if r.counter[0] != 0 {
		return
	}
	r.counter[1]++
	if r.counter[1] != 0 {
		return
	}
	r.counter[2]++
	if r.counter[2] != 0 {
		return
	}
	r.counter[3]++
}
