package basic_date

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

func TestRunHttp(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(500)
	for i := 0; i < 500; i++ {
		go func() {
			res, err := http.Get("http://localhost/?a=1")
			if err != nil {
				return
			}
			fmt.Println(res.StatusCode)
			wg.Done()
		}()
	}
	wg.Wait()
}
