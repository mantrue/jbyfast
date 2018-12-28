package basic_date

import (
	"fmt"
	"strings"
	"testing"
)

func TestInt64Max(t *testing.T) {
	var a uint64
	fmt.Println(a)
	t.Log("error")
}

func TestStringBuild(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("hello")
	builder.WriteString("123456789123456789123456789")
	fmt.Println(builder.String())
}

func TestBig(t *testing.T) {
	fmt.Println("test big...")
}

func TestSplit(t *testing.T) {
	fmt.Println(strings.Split("1|2", "|"))
}

func TestMap1(t *testing.T) {
	a := []interface{}{1}
	b := []interface{}{1}
	a = append(a, b...)
	fmt.Println(a)
}

func TestSlice(t *testing.T) {
	var slice = make([]int, 0, 3)
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice))
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Println(len(slice), cap(slice))
}
