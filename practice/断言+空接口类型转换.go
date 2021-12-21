package main

import "fmt"

//空接口的应用
//类型断言

//func assgin(arg interface{})  {
//  str, ok := arg.(string)
//  if !ok{
//      fmt.Printf("类型断言错误！\n")
//  }else {
//      fmt.Printf("恭喜你！猜对了；当前字符串内容为：%s\n", str)
//  }
//}

func assgin1(arg interface{}) {
	fmt.Printf("你输入的内容类型为：%T,", arg)
	switch t := arg.(type) {
	case string:
		fmt.Printf("内容为：%s\n", t)
	case int:
		fmt.Printf("内容为：%d\n", t)
	case bool:
		fmt.Printf("内容为：")
		fmt.Println(t)
	default:
		fmt.Println("***********")
		fmt.Printf("您输出的类型为：%#v,", arg)
	}
}

func main() {

	//assgin(666)
	//var str string
	var arg interface{}
	v, ok := arg.(string)
	fmt.Println(v, ok)

	//str := "hello Golang"
	assgin1(v)
	//fmt.Println(assgin2)

	//类型断言
	/*	v, ok := username.(string)
		if ok==true {
			c.String(200, "用户列表--"+v)
		} else {
			c.String(200, "用户列表--获取用户失败")
		}
	*/
}
