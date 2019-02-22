package go_testing_examples

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	var m sync.Mutex
	var count uint64
	for i := 0; i < b.N; i++ {
		m.Lock()
		count++
		m.Unlock()
	}
	fmt.Println(count)
}
