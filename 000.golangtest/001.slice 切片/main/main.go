package main

import (
	"fmt"
)

/*
切片与数组类似，可以理解为动态的数组，切片是局域数组实现的，他的底层是数组。对任意数组进行分隔
就可以得到切片
*/
func main() {
	//1.定义一个数组
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//2.对数组 arr 进行分隔得到一个切片
	sli := arr[3:6]
	sli2 := arr[:7]
	fmt.Printf("sli 的数据类型是%T\n数据内容是%d\n", sli, sli)
	fmt.Printf("sli 的内存地址是%p\n", sli)
	fmt.Printf("sli2 的内存地址是%p\n", sli2)
	fmt.Printf("arry 的内存地址是%p\n", &arr)
	fmt.Println("数组是sli 的长度是\n", len(sli))
	//3.切片元素的修改
	sli[1] = 33
	fmt.Printf("sli 的数据类型是%T\n数据内容是%d\n内存地址是%p\n", sli, sli, sli)
	fmt.Printf("arr 的数据类型是%T\n数据内容是%d\n内存地址是%p\n", arr, arr, &arr)
	sli3 := append(sli, 66, 99, 22)
	fmt.Printf("arr 的内存地址是:%p\nsli的内存地址是:%p\nsli3的内存地址是:%p\n", &arr, sli, sli3)
	sli4 := append(sli, 66, 99, 22, 44, 55, 66, 77, 88, 99, 100)
	fmt.Printf("arr 的内存地址是:%p\nsli的内存地址是:%p\nsli3的内存地址是:%p\nsli4 的内存地址是:%p\n", &arr, sli, sli3, sli4)
	fmt.Println(sli3)
	fmt.Println(sli)
	fmt.Println("--________________________")
	//数组的定义方法二
	slice1 := make([]int, 4, 4) //长度是 4，内存是 4
	//2.数组的默认值
	fmt.Println("slice1:", slice1)
	//3.向 slice1 中添加数据
	slice1[0] = 888
	slice1[1] = 999
	slice1[2] = 666
	slice1[3] = 777
	slice2 := append(slice1, 333, 555, 222, 111)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Printf("slice1的内存地址是:%p\nslice2的内存地址是:%p\n", slice1, slice2)
	fmt.Printf("slice1:%d\nslice2:%d\n", slice1, slice2)
	fmt.Println(len(slice2), cap(slice2))
	for i, v := range slice2 {
		fmt.Println(i, "------>", v)
	}

}
