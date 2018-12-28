package basic_date

import (
	"fmt"
	"testing"
)

type A struct {
	I int
	J int
}

func (a *A) Hello() {
	fmt.Println("helloworld")
}

func TestA(t *testing.T) {
	a := new(A)
	a.Hello()

	var a1 A
	a1.Hello()

	var a2 *A
	a2.Hello()
}
