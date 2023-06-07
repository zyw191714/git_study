package main

import "fmt"

func main() {
	/*
		课后作业 2：判断两条线是否平行
		提示：

		两点决定一条直线   斜率公式  k= （y1 - y2） / (x1 -x2 )
		两条线是否平行取决于两条线的斜率是否一样
	*/
	//1、定义环境变量
	var x1 float64
	fmt.Println("请输入x1 的值")
	fmt.Scanln(&x1)

	var x2 float64
	fmt.Println("请输入x2 的值")
	fmt.Scanln(&x2)
	fmt.Println(x2)

	var y1 float64
	fmt.Println("请输入y1 的值")
	fmt.Scanln(&y1)
	fmt.Println(y1)

	var y2 float64
	fmt.Println("请输入y2 的值")
	fmt.Scanln(&y2)
	fmt.Println(y2)
	//输出第一条线的k值
	var k1 float64 = (y2 - y1) / (x2 - x1)
	fmt.Println("k1的值是:", k1)

	//定义第二条线的数据
	var x3 float64
	fmt.Println("请输入x3 的值")
	fmt.Scanln(&x3)

	var x4 float64
	fmt.Println("请输入x4 的值")
	fmt.Scanln(&x4)
	fmt.Println(x4)

	var y3 float64
	fmt.Println("请输入y3 的值")
	fmt.Scanln(&y3)
	fmt.Println(y3)

	var y4 float64
	fmt.Println("请输入y4 的值")
	fmt.Scanln(&y4)git
	fmt.Println(y4)
	//输出第二条线的k值
	var k2 float64 = (y4 - y3) / (x4 - x3)
	fmt.Println("k1的值是:", k2)
	// 判断两条直线是否平行
	if k1 == k2 {
		fmt.Println("两条直线平行")
	} else {
		fmt.Println("两条直线不平行")
	}
}
