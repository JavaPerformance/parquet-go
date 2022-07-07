// Package wyhash implements a hashing algorithm derived from the Go runtime's
// internal hashing fallback, which uses a variation of the wyhash algorithm.
package wyhash

import (
	"encoding/binary"
	"math/bits"
)

const (
	m1 = 0xa0761d6478bd642f
	m2 = 0xe7037ed1a0b428db
	m3 = 0x8ebc6af09c88c6e3
	m4 = 0x589965cc75374cc3
	m5 = 0x1d8e4e27c47d124f
)

func mix(a, b uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	return hi ^ lo
}

func Hash32(value uint32, seed uintptr) uintptr {
	for {
		h := uintptr(mix(m5^4, mix(uint64(value)^m2, uint64(value)^uint64(seed)^m1)))
		if h != 0 {
			return h
		}
		seed++
	}
}

func Hash64(value uint64, seed uintptr) uintptr {
	for {
		h := uintptr(mix(m5^8, mix(value^m2, value^uint64(seed)^m1)))
		if h != 0 {
			return h
		}
		seed++
	}
}

func Hash128(value [16]byte, seed uintptr) uintptr {
	for {
		a := binary.LittleEndian.Uint64(value[:8])
		b := binary.LittleEndian.Uint64(value[8:])
		h := uintptr(mix(m5^16, mix(a^m2, b^uint64(seed)^m1)))
		if h != 0 {
			return h
		}
		seed++
	}
}
