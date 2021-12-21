package main

import (
	"fmt"
)

func main() {
	//1.类型和表达式
	//2.常量，一般作为全局常量，程序中不要太多全局常量，常量的值是在程序编译的时候确定的，之后不可再变。
	var name = true
	var name1 = "张总武功天下第一"
	fmt.Println(name)
	fmt.Println(name1)

	const LENGTH = 100
	const s1 = "test"
	const s2 string = "test"
	fmt.Println(s1)
	fmt.Println(s2)

	//iota：声明常量时可以使用常量生成器iota。iota可以通过枚举创建一系列相关的值，而且不需要明确定义类型。iota每次从0开始取值，逐次加1。
	type data int
	const (
		Zero data = iota
		One
		Two
		Three
		Four
	)
	fmt.Println(Zero)
	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)
	fmt.Println(Four)
	fmt.Println("")

	//灵活运用iota通过唯一运算符达到乘以2的幂的效果
	const (
		p2_0 = 1 << iota
		p2_1
		p2_2
		p2_3
	)
	fmt.Println(p2_0)
	fmt.Println(p2_1)
	fmt.Println(p2_2)
	fmt.Println(p2_3)
}
