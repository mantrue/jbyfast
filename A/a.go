package A

import (
	"A/B"
	"fmt"
)

var NameA string = "yesyes"

func init() {
	fmt.Println("AAAAAAAAAAAAAAAA")
	fmt.Println(B.FirstB)
	fmt.Println(B.Mystruct{100})
}
