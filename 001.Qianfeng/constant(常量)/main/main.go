package main

/*
常量的定义方式
	const 常量的名字（英文大写字母） [数据类型] = value
概念：同变量定义方式类似，程序执行过程中数值不能改变
变量和常量的区别：变量定义之后不使用会报错，但是常量不会
*/
import "fmt"

func main() {
	const BAIDU = "https://www.baidu.com"
	const PI = 3.14
	fmt.Println(BAIDU)
	fmt.Println(PI)
	//BAIDU = "http://www.baidu.com"  //cannot assign to BAIDU
	// 定义一组常量
	const NUM1, NAME1, SEX1 = 28, "张艳文", "男"
	fmt.Println(NUM1, NAME1, SEX1)
	const (
		NU = 1
		NL = 2
	)
	fmt.Println(NU, NL)
	// 在一组常量中，如果某个常量没有初始值，默认和上一行一致
	const (
		a int = 100
		b
	)
	fmt.Printf("%T,%d", a, a)
	fmt.Printf("%T,%d", b, b)
	// 枚举类型：使用常量组作为枚举类型，一组相关值的数据
	const (
		SPR = 0
		SUM = 1
		AUT = 2
		WIN = 3
	)
}
