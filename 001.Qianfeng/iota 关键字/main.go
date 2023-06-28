package main

import "fmt"

func main() {
	/* iota：特殊的常量，可以被编译器自动修改的常量，每定义一个 const，iota 的初始值为 0
	每当定义一个常量，就会自动累加 1，直到下一个常量的出现
	*/
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	const (
		d = iota
		e
	)
	fmt.Println(d, e)

}
