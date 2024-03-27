//go:build s390x

package deprecated

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"math/bits"
	"strings"
)

// Int96 is an implementation of the deprecated INT96 parquet type.
type Int96 [3]uint32

// Int32ToInt96 converts a int32 value to a Int96.
func Int32ToInt96(value int32) (i96 Int96) {
	fmt.Print("Int32ToInt96\n")
	if value < 0 {
		i96[0] = 0xFFFFFFFF // was 2
		i96[1] = 0xFFFFFFFF //was 1
	}
	i96[2] = uint32(value) // was 0

	PrintInt32BitPattern(i96[0])
	fmt.Print(" - ")
	PrintInt32BitPattern(i96[1])
	fmt.Print(" - ")
	PrintInt32BitPattern(i96[2])
	fmt.Print("\n")

	return
}

func PrintInt96BitPattern(i96 Int96) {

	PrintInt32BitPattern(i96[0])
	fmt.Print(" - ")
	PrintInt32BitPattern(i96[1])
	fmt.Print(" - ")
	PrintInt32BitPattern(i96[2])
	fmt.Print("\n")

}

func PrintInt32BitPattern(n uint32) {
	bits := make([]string, 32)

	for i := 0; i < 32; i++ {
		if n&(1<<(31-i)) != 0 {
			bits[i] = "1"
		} else {
			bits[i] = "0"
		}
	}

	// Insert spaces between every 8 bits (1 byte)
	for i := 8; i < 32; i += 9 {
		bits = append(bits[:i], append([]string{" "}, bits[i:]...)...)
	}

	// Join all bits into a single string
	bitPattern := strings.Join(bits, "")

	fmt.Print(bitPattern)
}

// Int64ToInt96 converts a int64 value to Int96.
func Int64ToInt96(value int64) (i96 Int96) {
	fmt.Print("Int64ToInt96\n")
	if value < 0 {
		i96[0] = 0xFFFFFFFF // was 2
	}
	i96[1] = uint32(value >> 32) // was 1
	i96[2] = uint32(value)       // was 0
	return
}

// IsZero returns true if i is the zero-value.
func (i Int96) IsZero() bool { return i == Int96{} }

// Negative returns true if i is a negative value.
func (i Int96) Negative() bool {
	fmt.Print("Int96ToBytes\n")
	return (i[0] >> 31) != 0
}

// Less returns true if i < j.
//
// The method implements a signed comparison between the two operands.
func (i Int96) Less(j Int96) bool {
	fmt.Print("Int96ToBytes\n")
	if i.Negative() {
		if !j.Negative() {
			return true
		}
	} else {
		if j.Negative() {
			return false
		}
	}
	for k := 2; k >= 0; k-- {
		a, b := i[k], j[k]
		switch {
		case a < b:
			return true
		case a > b:
			return false
		}
	}
	return false
}

// Int converts i to a big.Int representation.
func (i Int96) Int() *big.Int {
	fmt.Print("Int96ToBytes\n")
	z := new(big.Int)
	z.Or(z, big.NewInt(int64(i[0])<<32|int64(i[1])))
	z.Lsh(z, 32)
	z.Or(z, big.NewInt(int64(i[2])))
	return z
}

// Int32 converts i to a int32, potentially truncating the value.
func (i Int96) Int32() int32 {
	fmt.Print("Int32\n")
	return int32(i[2])
}

// Int64 converts i to a int64, potentially truncating the value.
func (i Int96) Int64() int64 {
	fmt.Print("Int64\n")
	return int64(i[1])<<32 | int64(i[2])
}

// String returns a string representation of i.
func (i Int96) String() string {
	fmt.Print("String\n")
	return i.Int().String()
}

// Len returns the minimum length in bits required to store the value of i.
func (i Int96) Len() int {
	fmt.Print("Len\n")
	switch {
	case i[0] != 0:
		return 64 + bits.Len32(i[2])
	case i[1] != 0:
		return 32 + bits.Len32(i[1])
	default:
		return bits.Len32(i[2])
	}
}

func Int96ToBytes(data []Int96) []byte {
	fmt.Print("Int96ToBytes\n")

	for i := 0; i < len(data); i++ {
		i96 := data[i]
		PrintInt32BitPattern(i96[0])
		fmt.Print(" - ")
		PrintInt32BitPattern(i96[1])
		fmt.Print(" - ")
		PrintInt32BitPattern(i96[2])
		fmt.Print("\n")
	}

	result := make([]byte, 0, len(data)*12) // Pre-allocate for efficiency

	for _, i96 := range data {
		for i := 2; i >= 0; i-- { // Iterate in reverse
			buf := make([]byte, 4)
			binary.LittleEndian.PutUint32(buf, uint32(i96[i]))
			result = append(result, buf...)
		}
	}

	return result
}

// BytesToInt96 converts the byte slice passed as argument to a slice of Int96
// sharing the same backing array.
//
// When the number of bytes in the input is not a multiple of 12, the function
// truncates it in the returned slice.

func BytesToInt96(data []byte) []Int96 {
	fmt.Print("BytesToInt96\n")
	if len(data)%12 != 0 {
		// Handle potential error if input data length is not divisible by 12
		return nil
	}

	result := make([]Int96, len(data)/12)

	for i := 0; i < len(result); i++ {
		start := i * 12
		for j := 2; j >= 0; j-- {
			value := binary.LittleEndian.Uint32(data[start : start+4]) // Big endian conversion
			result[i][j] = uint32(value)
			start += 4
		}
	}

	return result
}

func MaxLenInt96(data []Int96) int {
	fmt.Print("MaxLenInt96\n")
	max := 0
	for i := range data {
		n := data[i].Len()
		if n > max {
			max = n
		}
	}
	return max
}

func MinInt96(data []Int96) (min Int96) {
	fmt.Print("MinInt96\n")
	if len(data) > 0 {
		min = data[0]
		for _, v := range data[1:] {
			if v.Less(min) {
				min = v
			}
		}
	}
	return min
}

func MaxInt96(data []Int96) (max Int96) {
	fmt.Print("MaxInt96\n")
	if len(data) > 0 {
		max = data[0]
		for _, v := range data[1:] {
			if max.Less(v) {
				max = v
			}
		}
	}
	return max
}

func MinMaxInt96(data []Int96) (min, max Int96) {
	fmt.Print("MinMaxInt96\n")
	if len(data) > 0 {
		min = data[0]
		max = data[0]
		for _, v := range data[1:] {
			if v.Less(min) {
				min = v
			}
			if max.Less(v) {
				max = v
			}
		}
	}
	return min, max
}

func OrderOfInt96(data []Int96) int {
	fmt.Print("OrderOfInt96\n")
	if len(data) > 1 {
		if int96AreInAscendingOrder(data) {
			return +1
		}
		if int96AreInDescendingOrder(data) {
			return -1
		}
	}
	return 0
}

func int96AreInAscendingOrder(data []Int96) bool {
	fmt.Print("int96AreInAscendingOrder\n")
	for i := len(data) - 1; i > 0; i-- {
		if data[i].Less(data[i-1]) {
			return false
		}
	}
	return true
}

func int96AreInDescendingOrder(data []Int96) bool {
	fmt.Print("int96AreInDescendingOrder\n")
	for i := len(data) - 1; i > 0; i-- {
		if data[i-1].Less(data[i]) {
			return false
		}
	}
	return true
}
