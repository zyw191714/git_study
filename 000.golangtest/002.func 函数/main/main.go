package main

import (
	"errors"
	"fmt"
)

// 函数返回值命名
func main() {
	fmt.Println(sum(6, 7))
	result, err := sum(7, 7)
	fmt.Println(result, err)
}
func sum(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a or b 不能是负数")
	}
	sum = a + b
	err = nil
	return sum, nil
}
