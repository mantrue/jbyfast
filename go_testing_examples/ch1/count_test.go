package ch1

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCounts(t *testing.T) {
	//test()
	//test1()
	test2()
}

func test2() {
	counts := make(map[string]int)
	f, err := os.Open("E:/goroot/src/go_testing_examples/ch1/demo.ini")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewBuffer(data)

	for {
		line, err := reader.ReadString('\n')
		if line != "" && len(line) > 0 {
			counts[line]++
		}
		if err != nil {
			break
		}
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Println(k, "=======", v)
		}
	}
}

func test1() {
	counts := make(map[string]int)
	f, err := os.Open("E:/goroot/src/go_testing_examples/ch1/demo.ini")
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		if string(line) != "" && len(line) > 0 {
			counts[string(line)]++
		}
		if err != nil {
			break
		}
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Println(k, "=======", v)
		}
	}
}

func test() {
	counts := make(map[string]int)
	f, err := os.Open("E:/goroot/src/go_testing_examples/ch1/demo.ini")
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(data), "\r\n") {
		counts[line]++
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Println(k, "=============", v)
		}
	}
}
