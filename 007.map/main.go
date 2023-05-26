package main

import "fmt"

// map 键值对存在
func main() {
	var m1 map[string]int
	m2 := map[string]int{}
	m3 := map[string]int{"张三": 99, "李四": 88, "王二麻子": 77}
	// 定义一个空的map,map 必须实例化才可以写入
	var m4 map[int]int = nil
	fmt.Println("m4", m4)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	// 判断我们收集到的数据是否为真
	laoLiu, ok := m3["赵老六"]
	fmt.Println(laoLiu, ">>>>>>", ok)
	// 向map中增加元素
	m3["赵老六"] = 66
	zhaoLaoLiu, ok := m3["赵老六"]
	fmt.Println(zhaoLaoLiu, ">>>>>>", ok)
	fmt.Println(m3)
	// 删除map中的元素
	delete(m3, "王二麻子")
	fmt.Println(m3)
	// 修改map中的元素
	m3["张三"] = 78
	fmt.Println(m3)
	//遍历map,map函数没有固定的下标，所以遍历的时候区别于其他函数，没有 i++ 这种写法
	for k, v := range m3 {
		fmt.Println(k, v)
	}
}
