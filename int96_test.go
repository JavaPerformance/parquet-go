package parquet

import (
	"github.com/parquet-go/parquet-go/deprecated"
	"github.com/parquet-go/parquet-go/encoding/rle"
	"testing"
)

func TestInt96(t *testing.T) {

	t.Log("**********************************************************")

	i96a := deprecated.Int32ToInt96(1)
	i96b := deprecated.Int32ToInt96(2)

	t.Logf("i96 = %v", i96a)
	t.Logf("i96 = %v", i96b)

	int96Array := make([]deprecated.Int96, 2)

	int96Array[0] = i96a
	int96Array[1] = i96b

	result := deprecated.Int96ToBytes(int96Array)

	rle.PrintByteArrayBitPattern(result)

	t.Log("**********************************************************")

}
