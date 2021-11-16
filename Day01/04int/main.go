package main

import "fmt"

//整型

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) //把十进制数转化为二进制数
	fmt.Printf("%o\n", i1) //把十进制数转化为八进制数       一般给文件设置权限用八
	fmt.Printf("%x\n", i1) //把十进制数转化为十六进制数
	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0xf1234567
	fmt.Printf("%d\n", i3)
	//查看变量的类型
	fmt.Printf("%T\n", i3)

	//声明int8类型的变量
	i4 := int8(9) //明确指定int8 , 否则默认int
	fmt.Printf("%T\n", i4)
}
