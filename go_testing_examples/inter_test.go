package go_testing_examples

import (
	"bytes"
	"fmt"
	"testing"
)

var (
	i   = 5
	str = "ABC"
)

type Persion struct {
	name string
	age  int
}

type Any interface{}

func TestInter(t *testing.T) {
	var val Any
	val = 5
	val = str
	fmt.Println(val)

	pers1 := new(Persion)

	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)

	val = "dage"

	if v, ok := val.(string); ok {
		fmt.Println(v, "===============", ok)
	}

	switch t := val.(type) {

	case *Persion:
		fmt.Println("this is struct", t.name+"oks")
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %s\n", t+"oks")
	default:
		fmt.Printf("Unexpected type %T", t)

	}

	f := test()
	fmt.Println("1:", f())
	fmt.Println("2:", f())
	fmt.Println("3:", f())

	var my map[string][]string
	my = map[string][]string{"oks": []string{"dalao"}}
	my["two"] = append(my["two"], "oks")
	fmt.Println(my)

	buffer := [1000]byte{}
	fmt.Println(buffer)

	s := buffer[:0]
	fmt.Println(s)
	var buffer2 bytes.Buffer
	for i := 0; i < 10; i++ {
		buffer2.WriteString("1")
	}
	fmt.Println(buffer2.String())
}

func test() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
