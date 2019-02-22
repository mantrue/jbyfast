package go_testing_examples

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	dict := make(map[string]int)
	dict["小张"] = 28
	fmt.Println(dict)

	dict = map[string]int{"小李": 29}
	fmt.Println(dict)

	age, exists := dict["小王"]
	fmt.Printf("is exists:%v, age:%d\n", exists, age)

	delete(dict, "小张")
	fmt.Println(dict)

	for key, val := range dict {
		fmt.Println(key, val)
	}

	dict = map[string]int{"小五": 20, "小三": 23}
	modify(dict)
	fmt.Println(dict["小三"])
}

func modify(dict map[string]int) {
	dict["小三"] = 25
}
