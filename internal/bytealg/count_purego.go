//go:build purego || !amd64

package bytealg

import (
	"bytes"
	"fmt"
)

func Count(data []byte, value byte) int {

	count := bytes.Count(data, []byte{value})
	fmt.Printf("Count %d\n", count)
	return count
	//return bytes.Count(data, []byte{value})
}
