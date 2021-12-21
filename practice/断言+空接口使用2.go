package main

import "fmt"

func typeCast(x interface{}) {
	y, ok := x.(int) // 类型转换
	//fmt.Printf("q is ?: %#v",q)
	fmt.Println(y, ok) // 0 false

	z, ok1 := x.(string) // 类型转换
	fmt.Println(z, ok1)  // darren true
}

func switch_type_test(x interface{}) {
	// 1.不能直接使用x.(type)
	// fmt.Println(x.(type))  // 编译失败: use of .(type) outside type switch

	// 2.正确的x.(type)使用方式: 与switch...case...
	switch x.(type) {
	case string:
		fmt.Printf("x is string: %s\n", x)
		fmt.Printf("x is string: %#v", x)
		fmt.Println("")

	case int:
		fmt.Printf("x is int: %d\n", x)
		fmt.Printf("x is int: %#v", x)
		fmt.Println("")

	}
}
func main() {
	s := "darren"
	typeCast(s)

	fmt.Println("")

	q := 1234
	typeCast(q)

	switch_type_test(1)
	switch_type_test("Tom")
}
