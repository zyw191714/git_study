package main

import "fmt"

func main() {
	var money = 0
	switch money {
	case 20:
		fmt.Println("点个外卖")
	case 200:
		fmt.Println("去下馆子")
	case 2000:
		fmt.Println("到米其林体验体验")
	case 20000:
		fmt.Println("出国游玩一圈")
	default:
		fmt.Println("容我想想")
	}
}
