package main

import (
	"fmt"
	"strings"
)

//将字符串指定的后缀字符串移除
func main() {
	fmt.Println(strings.TrimSuffix("123aba1bab", "ab"))
	fmt.Println(strings.TrimSuffix("123abaabb", "ab"))

}

//func TrimSuffix(s, suffix string) string
