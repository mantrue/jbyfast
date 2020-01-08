package main

import "fmt"

type Animaler interface {
	Say() string
}

type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

func main() {
	c := Cat{}
	d := Dog{}
	var a Animaler

	a = c
	fmt.Println(a.Say())
	a = d
	fmt.Println(a.Say())
}
