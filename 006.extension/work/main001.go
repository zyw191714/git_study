package main

import (
	"fmt"
)

func main() {
	weekDays := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(weekDays)
	fmt.Println(len(weekDays))
	fmt.Println("请输入你想要查询的年份:")
	var years int
	//var yearsDay int
	var leapdays int = 0
	// 计算平年和闰年
	fmt.Scanln(&years)
	if years%4 == 0 && years%4 != 0 || years%4 == 0 {
		fmt.Println("今年是闰年，全年有366天")
		//yearsDay = 366
		leapdays = 1
	} else {
		fmt.Println("平年，365天")
		//yearsDay = 365

	}
	//定义月份表
	for month := 1; month <= 12; month++ {
		fmt.Println(month)
	}
	// 定义日期表

}
