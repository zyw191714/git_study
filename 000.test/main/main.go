package main

import "fmt"

func main() {
	weekdays := [...]string{"一", "二", "三", "四", "五", "六", "日"}
	month1 := [4][7]int{
		{1, 2, 3, 4, 5, 6, 7},
		{8, 9, 10, 11, 12, 13, 14},
		{15, 16, 17, 18, 19, 20, 21},
		{22, 23, 24, 25, 26, 27, 28},
	}
	month02 := [3]int{29, 30, 31}
	fmt.Println(weekdays)
	fmt.Println(month1[0])
	fmt.Println(month1[1])
	fmt.Println(month1[2])
	fmt.Println(month1[3])
	fmt.Println(month02)
	nums := make([]int, 0, 0)
	//nums := []int{}
	nums = append(nums, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(nums)
	fmt.Printf("%p\n", &month1)
	fmt.Printf("%p\n", nums)

}
