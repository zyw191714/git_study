package main

import "fmt"

func main() {
	// 定义月份，年份判断每个月的天数
	// 根据年份判断星期
	month := 0
	year := 0
	day := 0
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	fmt.Println("请输入年份")
	fmt.Scanln(&year)
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31
		fmt.Println(year, "年", month, "月份", day, "天")
	case 4, 6, 9, 11:
		day = 30
		fmt.Println(year, "年", month, "月份", day, "天")
	case 2:
		if year%400 == 0 || year%4 == 0 && year%100 != 0 {
			day = 29
			fmt.Println(year, "年", month, "月份", day, "天")
		} else {
			day = 28
			fmt.Println(year, "年", month, "月份", day, "天")
		}

	}

}
