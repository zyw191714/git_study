package main

import "fmt"

func main() {
	score := 0
	fmt.Println("请输入你的成绩")
	fmt.Scanln(&score)
	switch {
	case score >= 0 && score < 60:
		fmt.Println("不及格")
	case score >= 60 && score < 70:
		fmt.Println("及格")
	case score >= 70 && score < 80:
		fmt.Println("中等")
	case score >= 80 && score < 90:
		fmt.Println("良好")
	case score >= 90 && score <= 100:
		fmt.Println("优秀")
	default:
		fmt.Println("你输入的成绩有误")
	}
}
