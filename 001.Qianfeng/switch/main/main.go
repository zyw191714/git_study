package main

import "fmt"

func main() {
	/*
		使用 switch 语句写一个简单的计算器
	*/
	num1 := 0
	num2 := 0
	oper := "" // 定义一个空的变量
	fmt.Println("请输入一个整数:")
	fmt.Scanln(&num1)
	fmt.Println("请再次输入一个整数")
	fmt.Scanln(&num2)
	fmt.Println("请输入一个操作符:+,-,*,/")
	fmt.Scanln(&oper)
	switch oper {
	case "+":
		fmt.Println(num1, "+", num2, "=", num1+num2)
	case "-":
		fmt.Println(num1, "-", num2, "=", num1-num2)
	case "*":
		fmt.Println(num1, "*", num2, "=", num1*num2)
	case "/":
		fmt.Println(num1, "/", num2, "=", num1/num2)
	default:
		fmt.Println("你输入的操作符有误")
	}
	fmt.Println("main....over")
}
