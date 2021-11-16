package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	// \本来是具有特殊含义的，告诉程序 \ 是单纯的 \
	path := "C:\\Go\\bin"
	fmt.Println(path)

	//多行的字符串  反引号``
	s2 := `
		sdf 
	dsfasd
sdff

	`
	fmt.Println(s2)

	s3 := `C:\Go\bin` //反引号
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(s3))

	//字符串拼接
	name := "理想"
	world := "dsfsdf"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	//分隔
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(ss, "理想")) //true
	fmt.Println(strings.Contains(ss, "理性")) //false

	//前缀 后缀
	fmt.Println(strings.HasPrefix(ss, "理想")) //true
	fmt.Println(strings.HasSuffix(ss, "理想")) //false

	//判断子串在字符串中出现的位置
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))     //2
	fmt.Println(strings.Index(s4, "b"))     //1
	fmt.Println(strings.LastIndex(s4, "b")) //5   b最后一次出现的位置

	//join操作   拼接  strings.Join(a[]string, sep string)
	fmt.Println(strings.Join(ret, "+")) //C:+Go+bin

}
