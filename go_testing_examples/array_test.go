package go_testing_examples

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	var array1 [5]int
	fmt.Println(array1)
	var array2 [5]int
	array2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [5]int{1, 2, 3, 4, 5}
	array4 := [...]int{1, 2, 3, 4, 5}
	array5 := [5]int{0, 1, 0, 4, 0}
	array6 := [5]int{1: 1, 3: 4}

	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(array5)
	fmt.Println(array6)

	array8 := [5]*int{1: new(int), 3: new(int)}
	fmt.Println(array8)

	array9 := [5]int{1: 2, 3: 4}
	modifyArray(array9)
}

func modifyArray(a [5]int) {
	a[1] = 3
	fmt.Println(a)
}
