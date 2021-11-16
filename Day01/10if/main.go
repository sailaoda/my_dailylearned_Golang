package main

import "fmt"

//if 条件判断

func main() {
	age := 19 //作用域
	if age > 18 {
		fmt.Println("欢迎来到华科！")
	} else {
		fmt.Println("快去写作业")
	}

	// 多个判断条件
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好读书")
	}
}
