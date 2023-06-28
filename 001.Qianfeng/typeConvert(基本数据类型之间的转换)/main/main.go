package main

import "fmt"

/*
数据类型转换 type convert
兼容类型才可以转换
兼容类型之间才可以转换
常数：在需要的时候会自动转型
变量：需要手动转型
*/
func main() {
	var a int = 8
	var b int16
	// 将 a 转换成 int16 类型
	b = int16(a)
	fmt.Println(b)
	//var c bool
	//a = int(c)   cannot convert c (variable of type bool) to type int
}
