package main

import "fmt"

//用户结构体，用来声明对象，不赋值则为类型的初始值
type User struct {
	age     int
	name    string
	sex     string
	address string
	phone   string
	hobby   []string
}

func main() {
	//初始化对象
	var user1 User
	user1.age = 25
	user1.name = "Zying_Luo"
	user1.sex = "male"
	user1.address = "广东珠海香洲区"
	user1.phone = "15323616688"
	user1.hobby = []string{"编程", "游戏", "斗破"}

	fmt.Print(user1)
	fmt.Println()
	fmt.Printf("%v来自%v喜欢%v\n", user1.name, user1.address, user1.hobby)
	fmt.Printf("%T\n", user1)

	changeInfo(user1)

	//修改信息
	fmt.Printf("修改后的用户信息为%v\n", user1)
	changeInfo1(&user1)
	//修改信息
	fmt.Printf("真实修改后的用户信息为%v\n", user1)
}

//函数参数是拷贝，也就是副本，改变不了原来的类的值,d所以此处改变的20的值失败了
func changeInfo(user User) {
	user.sex = "female"
	user.age = 20
}

//如果需要改变值，要传指针，根据地址找到原来的值，再去修改！！！这个地址就是user1的地址!所以“&user1”，取了user1然后利用指针修改其属性
func changeInfo1(user *User) {
	//(*user).sex = "female"
	//指针是不能修改的，可以直接user去取,会根据指针取到对应的对象
	user.sex = "female"

	(*user).age = 18
}
