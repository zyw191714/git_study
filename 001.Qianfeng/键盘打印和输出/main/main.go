package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		格式化打印占位符
		%v 原样输出
		%T 打印类型
		%t bool 类型
		%s 字符串
		%f 浮点
		%d 十进制整数
		%b 二进制整数
		%o 8 进制
		%x，%X 十六进制
		%c 打印字符
		%p 打印地址(指针)

	*/
	a := 100
	b := 3.14
	c := true
	d := "hello"
	f := "A"
	fmt.Printf("变量 a 的类型是%T,表示的数值是%d\n", a, a)
	fmt.Printf("变量 b 的类型是%T,表示的数值是%f\n", b, b)
	fmt.Printf("变量 c 的类型是%T,表示的数值是%t\n", c, c)
	fmt.Printf("变量 d 的类型是%T,表示的数值是%s\n", d, d)
	fmt.Printf("变量 f 的类型是%T,表示的数值是%s \n", f, f)

	var x int
	var y float64
	fmt.Println("请输入一个整数x和一个浮点数y")
	fmt.Scanln(&x, &y) // 读取键盘的输入通过获取地址赋值给 x 和 y,阻塞式
	fmt.Printf("x 的数值是%d,y的数值是%.2f\n", x, y)
	fmt.Println("请输入一个整数x和一个浮点数y")
	fmt.Scanf("%d,%f\n", &x, &y)
	fmt.Printf("x 的数值是%d,y的数值是%.2f", x, y)

	fmt.Println("接下来是 bufio 输出部分————————————————————————")
	fmt.Println("请输入一个字符串")
	reader := bufio.NewReader(os.Stdin) // 标准键盘输入
	s1, _ := reader.ReadString('\n')    // 结束的标识符
	fmt.Println("读取到的数据是", s1)

}
