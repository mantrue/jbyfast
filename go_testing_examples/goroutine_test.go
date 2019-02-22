package go_testing_examples

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestRun(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 设置可用的物理处理器数量,CPU核数
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			fmt.Println("A:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			fmt.Println("B:", i)
		}
	}()
	wg.Wait()
}
