package main

//包声明
import "fmt"

// 导入包
func main() {
	var num int
	num = 100
	// 变量必须定义才能使用
	// 变量的类型和赋值必须一致
	//var name string = 100
	//fmt.Println(name)
	// 上面是错误示例:cannot use 100 (untyped int constant) as string value in variable declaration
	// 同一个作用域内变量不能重复声明(冲突)
	// 简短定义方式，左边的变量至少有一个是新的
	// 简短定义方式不能定义全局变量
	// 变量定义了就要使用
	fmt.Printf("num 的数值是:%d,地址是：%p\n", num, &num)
	num = 200
	fmt.Printf("num 的数值是:%d,地址是：%p\n", num, &num)

}
