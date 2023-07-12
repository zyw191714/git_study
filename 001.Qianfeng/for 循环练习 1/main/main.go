package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		水仙花数：
		水仙花数是一个三位数，每个位数的立方和，刚好等于这个数字本身，那么就叫水仙花数
	*/
	for i := 100; i < 1000; i++ {
		x := i / 100
		y := i / 10 % 10
		z := i % 10
		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)
		}
	}
	for a := 1; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				n := a*100 + b*10 + c*1
				if a*a*a+b*b*b+c*c*c == n {
					fmt.Println(n)

				}
			}
		}

	}
}
