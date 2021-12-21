package main

import (
	"fmt"
	"unicode/utf8"
)

//rune 数据类型 ==int32;byte==int8 中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
//string的底层是byte数组，即数组
func main() {
	var str = "中华文化博大精深 God"
	fmt.Println("len(str):", len(str))

	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))
	fmt.Println("rune类型来操作：", len([]rune(str)))
	fmt.Println("int32类型来操作：", len([]int32(str)))
	fmt.Printf("%#v\n", str)
	a := 123456789
	fmt.Println(a)
	fmt.Println(len(string(a)))

}
