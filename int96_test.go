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
