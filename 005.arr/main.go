package main

import "fmt"

/*
数组的定义
*/
func main() {
	var ages [5]int = [5]int{35, 36, 34, 43, 41}
	fmt.Println(ages)
	var ages2 = [3]int{31, 32, 33}
	fmt.Println(ages2)
	ages3 := [2]int{31, 23}
	fmt.Println(ages3)
	ages4 := [5]int{}
	fmt.Println(ages4)
	ages4[0] = 1
	ages4[1] = 11
	ages4[2] = 1111
	ages4[3] = 111
	ages4[4] = 11111
	fmt.Println(ages4)
	/*
	   遍历数组里面的元素
	*/
	for i := 0; i < len(ages4); i++ {
		fmt.Println(ages4[i])
	}
	for j := range ages {
		fmt.Println(ages[j])
	}
	for h, value := range ages2 {
		fmt.Printf("%d:%d\n", h, value)
	}
}
