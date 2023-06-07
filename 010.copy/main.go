package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	defer func() {
		finishTime := time.Now()
		fmt.Println(finishTime.Sub(startTime))
	}()
	// defer 函数定义的时候是由上到下进行定义的，运行是从下到上运行的
	arr2 := make([]string, 3, 4)
	// make([]string,3,4),里面第一个 3 是长度单位，第二个 4 是内存单位
	fmt.Println("len", len(arr2), "cap", cap(arr2))
	fmt.Println("default", arr2[0])
	fmt.Println("default", arr2[1])
	fmt.Println("default", arr2[2])
	arr3 := make([]string, 3, 4)
	copy(arr3, arr2)
	fmt.Println(arr3)

}
