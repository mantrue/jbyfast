package go_testing_examples

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

func TestBob(t *testing.T) {
	convert()
}

func convert() {
	stringSlice := []string{"通知中心", "perfect!"}

	buffer := &bytes.Buffer{}

	gob.NewEncoder(buffer).Encode(stringSlice)
	byteSlice := buffer.Bytes()
	fmt.Printf("%q\n", byteSlice)

	fmt.Println("---------------------------")

	backToStringSlice := []string{}
	gob.NewDecoder(buffer).Decode(&backToStringSlice)
	fmt.Printf("%v\n", backToStringSlice)
}
