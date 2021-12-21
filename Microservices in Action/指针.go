package main

import (
	"fmt"
)

func main() {
	//指针
	//1.
	fmt.Println("指针部分：")
	var x int       //x默认初始值为0
	p := &x         //p为整型指针，指向x
	fmt.Println(*p) //此时打印输出 0
	a := 10
	x = a
	fmt.Println(x)
	*p = 1         //类似于x=1
	fmt.Println(x) //此时打印输出1
	fmt.Println(a)

	//2.指针类型的默认初始值为nil
	m := 1
	selfPlusPointer(&m)       //n := &m
	fmt.Println(m)            //2
	fmt.Println(*selfPlus(1)) //2
}

func selfPlusPointer(n *int) {
	*n++ //就是指针n指向的地址m值++，所以是1+1=2；再深入点讲就是直接通过指针把值（m）修改了，就相当于把变量m直接改变一样，所以函数没有返回值
}

func selfPlus(n int) *int {
	t := n + 1
	return &t
}
