package main

import (
	"fmt"
)

func main() {
	// 多维数组的定义方式 定义了[3][3]string 是一个 3 * 3 的多维数组
	personInfo := [3][3]string{
		[3]string{"helen", "men", "22"},
		[3]string{"hello", "women", "33"},
		[3]string{"miss", "women", "22"},
	}
	for i, value := range personInfo {
		fmt.Printf("%v:%v\n", i, value)
		fmt.Println("___________________")
	}
	// 多维数组的第二种定义方式 可以不指定行数，指定列数进行定义，数组的定义形式不只是可以使用多个中括号[]进行指定
	newPersonInfo := [...][3]int{
		[3]int{1, 2, 3},
		[3]int{3, 4, 5},
		[3]int{5, 6, 7},
		[3]int{9, 9, 0},
	}
	// 多维数组的降维遍历方式
	for d1, d1val := range newPersonInfo {
		fmt.Println(d1, d1val)
		for d2, d2val := range d1val {
			fmt.Println(d2, d2val)
		}
	}
}
