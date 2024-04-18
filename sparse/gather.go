//go:build purego || !s390x

package sparse

import (
	"fmt"
	"unsafe"
)

func GatherInt32(dst []int32, src Int32Array) int {
	fmt.Println("gather.go GatherInt32")
	return GatherUint32(*(*[]uint32)(unsafe.Pointer(&dst)), src.Uint32Array())
}

func GatherInt64(dst []int64, src Int64Array) int {
	fmt.Println("gather.go GatherInt64")
	return GatherUint64(*(*[]uint64)(unsafe.Pointer(&dst)), src.Uint64Array())
}

func GatherFloat32(dst []float32, src Float32Array) int {
	fmt.Println("gather.go GatherFloat32")
	return GatherUint32(*(*[]uint32)(unsafe.Pointer(&dst)), src.Uint32Array())
}

func GatherFloat64(dst []float64, src Float64Array) int {
	fmt.Println("gather.go GatherFloat64")
	return GatherUint64(*(*[]uint64)(unsafe.Pointer(&dst)), src.Uint64Array())
}

func GatherBits(dst []byte, src Uint8Array) int {
	fmt.Println("gather.go GatherFloat64")
	return gatherBits(dst, src)
}

func GatherUint32(dst []uint32, src Uint32Array) int { return gather32(dst, src) }

func GatherUint64(dst []uint64, src Uint64Array) int { return gather64(dst, src) }

func GatherUint128(dst [][16]byte, src Uint128Array) int { return gather128(dst, src) }

func GatherString(dst []string, src StringArray) int {
	n := min(len(dst), src.Len())

	for i := range dst[:n] {
		dst[i] = src.Index(i)
	}

	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
