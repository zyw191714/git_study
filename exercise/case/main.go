package main

import "fmt"

func main() {
	var money float64 = 200
	var busy bool = false
	//定义两个变量，money和busy，通过输入不同的值来验证case语句返回的内容。
	switch {
	case money >= 0 && money <= 20:
		fmt.Println("点一个外卖")
		if busy {
			break
		}
		fmt.Println("再吃个零食")
	case money > 20 && money <= 200:
		fmt.Println("吃顿好的")
		if busy {
			break
		}
		fmt.Println("再吃个零食")
	default:
		fmt.Println("let me see")
	}
	fmt.Println("over")
}
