package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, "hello")
	}

}

/*
定义一个变量i，比较i，让变量进行累加，输出hello
运行逻辑描述：
首先定义一个变量 i = 0，条件是如果 i < 100，那么就打印hello，进入循环表达式 i++ ，当i = 1 的时候
进行判断，i++ 的结果依然小于100，然后继续print hello，直到 i = 100的时候结束循环
*/
