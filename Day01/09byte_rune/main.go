package main

import "fmt"

// byte和rune类型

// uint8类型  也叫 byte 型， 代表了ASCII 码的一个字符
// rune类型   代表一个UTF-8字符  rune类型实际上是一个int32
// Go 语言中为了处理非ASCII码类型的字符，定义了新的rune类型

func main() {
	s := "Hello华科"
	//len()是求得byte字节的数量
	n := len(s)
	fmt.Println(n)

	// for i := 0; i < len(s); i ++ {
	//	fmt.Println(s[i])
	//	fmt.Printf("%c\n", s[i])    //%c:字符
	//}

	// for _, c := range s {   //从字符串中拿出具体的字符
	// 	//fmt.Println(c)
	// 	fmt.Printf("%c\n", c)
	//}

	s2 := "白萝卜"             //  =>  '白'  '萝'  '卜'
	s3 := []rune(s2)        //把字符串强制转换成一个rune 切片
	s3[0] = '红'             //不是双引号"红" 表示字符串
	fmt.Println(string(s3)) //把rune 切片强制转换成字符串

	c1 := '红' //rune(int32)
	c2 := "红" //string
	fmt.Printf("c1:%T  c2:%T\n", c1, c2)

	//类型转换
	i := 10
	var f float64
	f = float64(i)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
