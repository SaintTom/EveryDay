package main

import "fmt"

// 空接口

func main() {
	var temp map[string]interface{} // 声明一个空接口类型变量
	temp = make(map[string]interface{}, 20)
	temp["name"] = "zhangsan"
	temp["weight"] = 65
	temp["school"] = true
	temp["hobel"] = [...]string{"打球", "跑步", "codeing"}
	fmt.Println(temp)
}
