package main

import "fmt"
func main(){
	a := []string{"I","am","stupid","and","weak"}
	// 定义一个不定长字符串数组a
	for i:=0;i<len(a);i++{
		fmt.Printf("i=%v\n",a[i])
	}
}