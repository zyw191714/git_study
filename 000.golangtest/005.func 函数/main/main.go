package main

import "fmt"

// 函数的数据类型
func main() {
	s1 := []int{1, 2, 3, 4}
	fmt.Println("函数调用前切片的数据是", s1)
	func2(s1)
	fmt.Println("函数调用后切片的数据是", s1)
	fmt.Println("------一条华丽的分割线---------")
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前，数组的数据：", arr1) //[1 2 3 4]
	fun1(arr1)
	fmt.Println("函数调用后，数组的数据：", arr1) //[1 2 3 4]

}
func func2(s2 []int) {
	fmt.Println("函数中，切片的数据是", s2)
	s2[0] = 100
	fmt.Println("函数中，切片的数据更改后", s2)
}
func fun1(arr2 [4]int) {
	fmt.Println("函数中，数组的数据：", arr2) //[1 2 3 4]
	arr2[0] = 100
	fmt.Println("函数中，数组的数据修改后：", arr2) //[100 2 3 4]
}
