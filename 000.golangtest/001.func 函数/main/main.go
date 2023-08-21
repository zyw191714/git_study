package main

import (
	"errors"
	"fmt"
)

/*
函数的组成：
	1.任何一个函数的定义都需要有一个 func关键字，用于声明一个函数
	2.后面需要跟函数的名字，符合 go 语言的命名规范即可
	3.函数名字后面的（）是不能省略的，（）里面可以定义函数使用的参数
	4.（）后面的函数可以有返回值，可以在大括号里面书写代码
*/
// 定义一个加法函数
func main() {
	sum, err := sumNum(1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
	//result, result1 := sumNum2(10, 11)
	//fmt.Println(result, result1)
}
func sumNum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("参数不能为负数")
	}
	return a + b, nil
}

//func sumNum2(A, B int) (int, int) {
//	return A + B, A - B
//}
