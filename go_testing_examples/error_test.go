package go_testing_examples

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := errors.New("whoops")
	fmt.Println(err)           // 简单的输出错误信息
	fmt.Printf("%+v\n\n", err) // 错误信息加堆栈
	fmt.Errorf("%s", "error")
}
