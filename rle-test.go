package parquet

import (
	"github.com/parquet-go/parquet-go/encoding/rle"
	"testing"
)

func TestEncodeDecodeInt32(t *testing.T) {

	t.Log("**********************************************************")

	rle.TestEncodeDecodeInt32(t)

	t.Log("**********************************************************")

}
