package main

import "fmt"

func main() {
	/*
		映射，映射是一个无序的 K-V键值对的集合，map 中所有的 key 必须具有相同
		的类型，value也一样，但是k 和 v 的类型可以不同，key 的类型必须支持==比较
		运算符，这样才能判断是否存在，并保证 key 的唯一性
	*/
	//1.定义一个 map
	nameAgeMap := make(map[string]int)
	nameAgeMap["张三"] = 20
	nameAgeMap["李斯"] = 21
	nameAgeMap["王武"] = 22
	nameAgeMap["赵六"] = 23
	nameAgeMap["陈琪"] = 24
	nameAgeMap["狗蛋"] = 25
	fmt.Println(nameAgeMap)
	fmt.Println(nameAgeMap)
	fmt.Println(nameAgeMap)
	fmt.Println(nameAgeMap["周天"])
	//3.map 的遍历
	for k, v := range nameAgeMap {
		fmt.Println(k, "------>", v)
	}
	delete(nameAgeMap, "狗蛋")
	//2.map的删除
	fmt.Println(nameAgeMap)
	fmt.Println(len(nameAgeMap))
	va, ok := nameAgeMap["张艳文"]
	if ok {
		fmt.Println(va)
	} else {
		fmt.Println("error")
	}

}
