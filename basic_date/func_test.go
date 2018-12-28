package basic_date

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func hello(a int) {
	fmt.Println("hello world")
}

func TestFunc(t *testing.T) {
	var f func(int)
	fty := reflect.TypeOf(f)
	ty := reflect.TypeOf(hello)

	fmt.Println(fty)
	fmt.Println(ty)
	fmt.Println(fty == ty)
}

func TestFunc2(t *testing.T) {
	f()
}

var i = 0

func f() {
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println(pc)
	fmt.Println(runtime.FuncForPC(pc))
	if i == 1 {
		panic("stop")
	}
	i++
	f2()
}

func f2() { // 记录下首次开启事务的函数
	pc, _, _, _ := runtime.Caller(0)
	fmt.Println(pc)
	fmt.Println(runtime.FuncForPC(pc))

	f()
}
