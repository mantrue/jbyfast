package go_testing_examples

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var myslice []string = []string{"a", "b", "c"}
	var index = 1
	myslice = append(myslice[:index], myslice[index+1:]...)

	for i, v := range myslice {
		fmt.Println(i, "=========", v)
	}

	slice1 := make([]int, 5)
	fmt.Printf("length:%d, cap:%d\n", len(slice1), cap(slice1))
	fmt.Println(slice1)

	slice2 := make([]int, 5, 10)
	fmt.Printf("length:%d, cap:%d\n", len(slice1), cap(slice2))
	fmt.Println(slice2)

	slice4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("length:%d, cap:%d\n", len(slice4), cap(slice4))

	slice5 := []int{5: 5}
	fmt.Printf("length:%d, cap:%d\n", len(slice5), cap(slice5))

	fmt.Println(slice4)
	fmt.Println(slice5)

	slice := []int{1, 2, 3, 4, 5}
	slice6 := slice[:]
	slice7 := slice[0:]
	slice8 := slice[:5]
	slice9 := slice[2:5]
	fmt.Println(slice6)
	fmt.Println(slice7)
	fmt.Println(slice8)
	fmt.Println(slice9)

	slice = []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]
	newSlice[1] = 10
	fmt.Println(slice)
	fmt.Println(newSlice)

	newSlice2 := slice[1:2:3]
	fmt.Println(newSlice2)
	fmt.Printf("length:%d, cap:%d\n", len(newSlice2), cap(newSlice2))

	fmt.Println("------------")

	slice = []int{1, 2, 3, 4, 5}
	newSlice = slice[1:3]

	newSlice = append(newSlice, 10)
	fmt.Println(newSlice)
	fmt.Println(slice)

	newSlice = append(newSlice, 10, 20)
	fmt.Println(newSlice)

	newSlice = append(newSlice, slice...)
	fmt.Println(newSlice)
	fmt.Println(slice)

	fmt.Println("----------")
}
