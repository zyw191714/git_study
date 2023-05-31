package main

import "fmt"

func main() {
	// 输入年份
	var year int
	fmt.Print("请输入年份：")
	fmt.Scanln(&year)

	// 计算这一年的天数和闰年数
	daysInYear := 365
	leapDays := 0
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		daysInYear = 366
		leapDays = 1
	}

	// 定义日历表
	calendar := make([][]string, 12)
	for i := 0; i < 12; i++ {
		calendar[i] = make([]string, 7)
	}

	// 初始化日历表
	for i := 0; i < 12; i++ {
		for j := 0; j < 7; j++ {
			calendar[i][j] = "  "
		}
	}

	// 填充日历表
	monthDays := []int{31, 28 + leapDays, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	weekDays := []string{"一", "二", "三", "四", "五", "六", "日"}
	month := 0
	day := 1
	weekDay := (year + (year-1)/4 - (year-1)/100 + (year-1)/400) % 7 // 计算这一年1月1日是星期几
	for i := 0; i < daysInYear; i++ {
		if i > 0 && i%monthDays[month] == 0 {
			month++
			day = 1
		}
		calendar[month][weekDay] = fmt.Sprintf("%2d", day)
		day++
		weekDay = (weekDay + 1) % 7
	}

	// 输出日历表
	fmt.Printf("%d年日历表\n", year)
	fmt.Println("一 二 三 四 五 六 日")
	for i := 0; i < 12; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf("%s ", calendar[i][j])
		}
		fmt.Println()
	}
}
