package main

import "fmt"
func main(){
	var a = []string{"I","am","stupid","and","weak"}
	fmt.Println("a数组的原始内容为",a)
	for i ,value := range a {
		fmt.Printf("key=%v;value=%v\n",i,value)
	}
	a[2]="smart"
	a[4]="strong"
	fmt.Println("修改后数组a的内容为",a)
}