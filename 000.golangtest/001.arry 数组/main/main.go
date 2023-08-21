package main

import "fmt"

func main() {
	//声明一个数组 数组名称 :=[数组长度]数组类型{数组内容(可以为空)}
	arry1 := [6]int{}
	arry2 := [6]string{}
	arry3 := [6]float32{}
	fmt.Println(arry1) // 打印输出没有定义内容时的默认值
	fmt.Println(arry2) // 打印输出没有定义内容时的默认值
	fmt.Println(arry3) // 打印输出没有定义内容时的默认值
	nums := [6]int{90, 91, 92, 93, 94, 95}
	fmt.Printf("%T\n", nums)
	fmt.Printf("数组 nums 的长度是%d，数组 nums 的容量是%d", len(nums), cap(nums))
	// 遍历数组 nums
	for i := 0; i <= len(nums); i++ {
		fmt.Println(i)
	}
	for i, v := range nums {
		// i,v是一对键值对，i 是 index，v 是对应的值。
		fmt.Println(i, v)
	}
	for _, v := range nums {
		// i,v是一对键值对，i 是 index，v 是对应的值。用"_",代表你想忽略不打印的值
		fmt.Println(v)
	}
	//修改数组中的数据
	nums2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("nums2 的长度是", len(nums2))
	fmt.Println("nums2：", nums2)
	nums2[9] = 10
	fmt.Println(nums2)
	nums3 := nums2
	fmt.Println(nums3)
	fmt.Printf("nums2 的内存地址是%p\n", &nums2)
	fmt.Printf("nums3 的内存地址是%p\n", &nums3)
	nums2[6] = 100
	fmt.Println(nums2)
	fmt.Println(nums3)

}
