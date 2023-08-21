package main

import "fmt"

// 可变参数，可变参数就是函数的参数数量是可变的
func main() {
	result := sum(6, 7, 8, 9, 22)
	fmt.Println(result)
	result1 := sum1(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(result1)
	//匿名函数
	sum := func(nums ...int) int {
		result := 0
		for _, i := range nums {
			result += i
		}
		return result
	}
	fmt.Println(sum(1, 2, 3, 4))
}
func sum(params ...int) int { // 可变参数类型实际上就是切片
	result := 0
	for _, i := range params {
		result += i
	}
	return result
}
func sum1(params ...int) int {
	result1 := 0
	for i := 0; i <= len(params); i++ {
		result1 += i
	}
	return result1
}
