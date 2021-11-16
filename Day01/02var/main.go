package main

import "fmt"

//函数外的每个语句都必须以关键字开始（var,const,func等）
//同一个作用域（{}）中不能重复声明同名的变量
//声明变量
var name string
var age2 int
var isOk bool

//批量声明
var (
	name2 string //不会报错
	age   int
	isOk2 bool
)

//匿名变量，用一个下划线_表示
func foo() (int, string) {
	return 10, "sailaoda"
}

func main() {
	//var name3 string       函数外部的声明的全局变量不使用不会报错只会警告，函数内部的变量必须使用
	name = "理想"
	age = 16
	isOk = true

	fmt.Print(isOk)
	fmt.Printf("name:%s \n", name)
	fmt.Println(age)
	fmt.Println(isOk2)

	//声明变量同时赋值
	var s1 string = "sailaoda"
	fmt.Println(s1)
	//类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	//简短变量声明  只能在函数中使用
	s3 := "哈哈哈"
	fmt.Println(s3)

	x, _ := foo() //匿名变量不占用空间，不会分配内存
	fmt.Println(x)
}
