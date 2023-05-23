/*
创建一个数组并进行反转
*/
package main

import "fmt"

func main() {
	numbers1 := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	nNumbers1 := [11]int{}
	//存在问题： 定义一维数组的长度必须指定，不然计算会失败
	n := len(numbers1)
	fmt.Println(n)
	for i := 0; i <= n/2; i += 1 {
		// 疑问？ 为什么定义 i <= n/2;
		fmt.Println(i)
		nNumbers1[i], nNumbers1[n-i-1] = numbers1[n-i-1], numbers1[i]
		fmt.Println(nNumbers1)
	}

}
