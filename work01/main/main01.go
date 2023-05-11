package main

import "fmt"
func main(){
	var a = []string{"I","am","stupid","and","weak"}
	for i:=0;i<len(a);i++{
		fmt.Printf("i=%v\n",a[i])
	}
}