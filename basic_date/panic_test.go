package basic_date

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	recoverPanic()
	myPanic()
}
func recoverPanic() {
	if p := recover(); p != nil {
		fmt.Println(p)
	}

}

func myPanic() error {
	panic("panic")
}
