package main

import (
	"fmt"
	"math"
)

//有问题啊？？？ 竟然不等于
func main() {
	//var mathYes bool
	a := math.MaxInt64
	b := math.MaxInt64 + 1
	c := math.MinInt64
	fmt.Println(a, b, c)
	if b == c {
		fmt.Println("math.MaxInt32+1=math.MinInt32")
		fmt.Println()
	}
}
