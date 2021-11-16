package main

import "fmt"

//常量
const pi = 3.1415926

//批量声明常量
const (
	statusOK = 200
	notFOUNT = 404
)

// 默认和上一行一样 n1 = 100 = n2 = n3
const (
	a1 = 100
	a2
	a3
)

// iota类似枚举
const (
	n1 = iota // 0
	n2 = iota // 1
	n3        // 2
	_         // 3
	n4        // 4
)

//插队
const (
	c1 = iota // 0
	c2 = 100
	c3 = iota // 2
	c4
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1 : 1    d2 : 2
	d3, d4 = iota + 1, iota + 2 //d3 : 2    d4 : 3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) //1 左移10 位  ； 10000000000    ： 1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3, n4)

	fmt.Println(c1, c2, c3, c4)

	fmt.Println(d1, d2, d3, d4)
}
