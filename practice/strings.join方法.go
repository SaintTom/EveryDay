package main

import (
	"fmt"
	"strings"
)

//golang字符串操作
func main() {
	//s := "hello world hello world"
	//str := "wo"
	var s = []string{"11", "22", "33"}

	//将s中的子串连接成一个单独的字符串，子串之间用str分隔。
	fmt.Println(s)
	fmt.Println("------------------------")
	ret := strings.Join(s, "|")
	fmt.Println(ret) //11|22|33
}
