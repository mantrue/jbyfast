package B

import (
	"fmt"
)

var FirstB string = "first B"

func init() {
	fmt.Println("first init...")
}

type Mystruct struct {
	Age int
}
