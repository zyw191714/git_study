package main

import "fmt"

func main() {
	// 定义一个切片
	a := []int{}
	fmt.Println(a)
	//向切片中追加元素
	a = append(a, 11, 22, 33, 44, 55, 66, 77, 88, 99, 00)
	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}
	fmt.Println("初始化", a)
	//删除切片中的元素
	//a = append(a[:2], a[3:]...)
	//fmt.Println("第一次删除", a)
	//a = append(a[:1], a[5:6]...)
	//fmt.Println("第二次删除", a)
	a = append(a[:1], a[:3]...)
	fmt.Println("第三次删除", a)
	// b是数组
	b := [0]int{}
	fmt.Println(b)
	fmt.Println("----------新生成的切片-----------")
	news := []int{11, 22, 33, 44, 55, 66}
	//news = append(news[:1], news[3:]...)
	news = append(news[:2], news[4:]...)
	fmt.Println(news)
}
