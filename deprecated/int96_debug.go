package deprecated

import (
	"fmt"
	"strings"
)

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

func PrintBitsWithSpaces(data []byte) {
	for _, byte := range data {
		fmt.Printf("%08b ", byte) // %08b formats as a zero-padded 8-bit binary number
	}
	fmt.Println() // Add a newline at the end
}

func PrintBitsWithSpacesWithLen(data []byte, len int) {
	for i := 0; i < len; i++ {
		fmt.Printf("%08b ", data[i]) // %08b formats as a zero-padded 8-bit binary number
	}
	fmt.Println() // Add a newline at the end
}
