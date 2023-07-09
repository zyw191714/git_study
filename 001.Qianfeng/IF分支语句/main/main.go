package main

import "fmt"

func main() {
	a := 100
	if a%10 == 0 {
		fmt.Println("这是个偶数")
	} else {
		fmt.Println("这不是一个偶数")
	}
	fmt.Println("main...over")

}
