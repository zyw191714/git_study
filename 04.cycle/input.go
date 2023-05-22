package main

import "fmt"

func main() {
	var name string
	fmt.Println("请输入你的名字")
	fmt.Scanln(&name)
	fmt.Scanf("%s", name)
	fmt.Println(name)

}
