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
}
