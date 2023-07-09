package main

/*
打印输出 100 以内，能被3 整除，不能被 5 整除的数字，每 5 个打印一行
*/
import "fmt"

func main() {
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Printf("%d\t", i)
			count++
			if count%5 == 0 {
				fmt.Println('\n')
			}
		}
	}
	fmt.Println()
	for x := 1; x < 10; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d x %d = %d\t", y, x, y*x)
		}
		fmt.Println()
	}
}
