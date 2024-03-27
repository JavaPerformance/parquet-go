//go:build !s390x

package parquet

import (
	"fmt"
	"github.com/parquet-go/parquet-go/deprecated"
	"github.com/parquet-go/parquet-go/encoding/rle"
	"testing"
)

func TestInt96(t *testing.T) {

	t.Log("**********************************************************")

	i96a := deprecated.Int32ToInt96(1)
	i96b := deprecated.Int32ToInt96(2)

	if i96a.Less(i96b) {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is greater than b")
	}

	if i96b.Less(i96a) {
		fmt.Println("b is less than a")
	} else {
		fmt.Println("b is greater than a")
	}

	t.Logf("i96 = %v", i96a)
	t.Logf("i96 = %v", i96b)

	int96Array := make([]deprecated.Int96, 2)

	int96Array[0] = i96a
	int96Array[1] = i96b

	t.Log("************************Int96ToBytes**********************************")

	result := deprecated.Int96ToBytes(int96Array)

	rle.PrintByteArrayBitPattern(result)

	t.Log("************************BytesToInt96**********************************")

	int96ArrayIn := deprecated.BytesToInt96(result)

	for i := 0; i < len(int96ArrayIn); i++ {
		i96 := int96ArrayIn[i]
		deprecated.PrintInt32BitPattern(i96[0])
		fmt.Print(" - ")
		deprecated.PrintInt32BitPattern(i96[1])
		fmt.Print(" - ")
		deprecated.PrintInt32BitPattern(i96[2])
		fmt.Print("\n")
	}

	t.Log("**********************************************************")

}

func TestInt96Len(t *testing.T) {
	testCases := []struct {
		i           deprecated.Int96
		expectedLen int
	}{
		{deprecated.Int96{0xFFFFFFFF, 0xFFFFFFFF, 1}, 65},
		{deprecated.Int96{1 << 31, 1, 0}, 33},
		{deprecated.Int96{0, 1 << 31, 0}, 64},
		{deprecated.Int96{1 << 31, 0, 0}, 32},
		{deprecated.Int96{123, 0, 0}, 7},
		{deprecated.Int96{65535, 0, 0}, 16},
		{deprecated.Int96{0, 0, 0}, 0},
		{deprecated.Int96{0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF}, 96}, // Fully set
	}

	for _, tc := range testCases {
		result := tc.i.Len()
		//		t.Logf("PASS Len() for %v to be %d, got %d", tc.i, tc.expectedLen, result)
		if result != tc.expectedLen {
			t.Errorf("Expected Len() for %v to be %d, got %d", tc.i, tc.expectedLen, result)
		}
	}
}
