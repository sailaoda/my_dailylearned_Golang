package main

import "fmt"

//for 循环

func main() {
	//基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 5
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	m := 8
	for m < 10 {
		fmt.Println(m)
		m++
	}

	//for range 键值循环
	s := "Hello华科"
	for i, v := range s {
		fmt.Printf("%d  %c\n", i, v)
	}
}
