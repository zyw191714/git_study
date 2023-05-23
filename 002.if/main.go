package main

import "fmt"

func main() {
	var money int = 35
	if money <= 20 {
		fmt.Println("就这，点个外卖都不够")
	} else if money <= 40 {
		fmt.Println("可以勉强吃顿外卖了")
	} else {
		fmt.Println("等会")
	}
}
