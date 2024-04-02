package parquet

import (
	"fmt"
	"github.com/parquet-go/parquet-go/internal/bytealg"
)

func countLevelsEqual(levels []byte, value byte) int {
	return bytealg.Count(levels, value)
}

func countLevelsNotEqual(levels []byte, value byte) int {
	lne := len(levels) - countLevelsEqual(levels, value)
	fmt.Printf("countLevelsNotEqual: %d %d\n", lne, value)

	//	buf := make([]byte, 1<<9) // Adjust buffer size as needed
	//	runtime.Stack(buf, false)
	// fmt.Println(string(buf))

	return lne
	//return len(levels) - countLevelsEqual(levels, value)
}

func appendLevel(levels []byte, value byte, count int) []byte {
	i := len(levels)
	n := len(levels) + count

	if cap(levels) < n {
		newLevels := make([]byte, n, 2*n)
		copy(newLevels, levels)
		levels = newLevels
	} else {
		levels = levels[:n]
	}

	bytealg.Broadcast(levels[i:], value)
	return levels
}
