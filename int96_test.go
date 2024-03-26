package parquet

import (
	"github.com/parquet-go/parquet-go/deprecated"
	"github.com/parquet-go/parquet-go/encoding/rle"
	"testing"
)

func TestInt96(t *testing.T) {

	t.Log("**********************************************************")

	i96 := deprecated.Int32ToInt96(1)

	t.Logf("i96 = %v", i96)

	int96Array := make([]deprecated.Int96, 1)

	int96Array[0] = i96

	result := deprecated.Int96ToBytes(int96Array)

	rle.PrintByteArrayBitPattern(result)

	t.Log("**********************************************************")

}
