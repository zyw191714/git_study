package main

import "fmt"

func main() {
	// 定义一个变量
	var num1 int = 8
	num2 := 8
	fmt.Println(num1, num2)
	/*
			定义一个数组
				1.数组的定义：数组(Array)存放的是固定长度、相同类型的数据，而且数组中的元素在内存中都是连续存放的。
		          它所存放的数据类型没有限制，可以是整型、字符串，甚至是自定义类型。
	*/
	fmt.Println("下面是 arry")
	arry1 := []int{1, 2, 3, 4, 5, 6}
	var arry2 [4]int
	arry2[0] = 3 // 向数组 arry2 中添加数据
	fmt.Println(arry1, arry2)
	/*
		定义一个切片
			1.切片的定义
			切片(Slice)与数组类似，可以把它理解为动态数组。切片是基于数组实现的，它的底层就是数组。对数组任意分隔
			，就可以得到切片。
	*/
	fmt.Println("下面是 slice")
	arry4 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := arry4[3:]
	fmt.Println(slice1)
	slice1[1] = 99 // 修改切片中的元素
	fmt.Println(arry4)
	fmt.Println(slice1)
	slice2 := append(slice1, 100, 101)
	fmt.Println("slice2", slice2)
	fmt.Println("arry4", arry4)
	fmt.Println("slice1", slice1)
	slice3 := make([]string, 4)
	fmt.Println(slice3)
	fmt.Println("下面是切片")

	/*
				定义一个 map
					1.在Go语言中，映射(Map)是一个无序的K-V（键值对）集合，结构为map[K]V。其中K对应Key，V对应Value。
		             map中所有的Key必须具有相同的类型，Value也一样，但Key和Value的类型可以不同。此外，Key的类型必须
		             支持==比较运算符，这样才可以判断它是否存在，并保证Key的唯一性。
	*/
	nameAgent := make(map[string]int)
	nameAgent["张三"] = 20
	fmt.Println(nameAgent)
	ageSex := map[string]int{"男": 22}
	fmt.Println(ageSex)
}
