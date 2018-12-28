package basic_date

import "testing"

func Test_error(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()
	go func() {
		panic("panic error!")
	}()
}
