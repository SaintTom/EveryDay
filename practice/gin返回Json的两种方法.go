package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回Json数据的两种方法：
func main() {
	//1.路由引擎返回
	//for i:=0; i<5; i++{
	//
	//}
	type HH map[int]interface{}
	r := gin.Default()
	r.GET("/gin", func(c *gin.Context) {
		c.JSON(http.StatusOK, HH{
			1: "张总", //复合文字后面要加英文的逗号
			2: "张总武功天下第一",
			3: 28,
			4: "male",
		})

	})

	//2.结构体返回，灵活使用tag对结构体定制化操作
	type zz struct {
		Name    string
		Age     int
		Message string
		Sex     string
	}
	r.GET("/struct", func(c *gin.Context) {
		data := zz{
			Name:    "张总",
			Age:     29,
			Message: "张总武功天下第一",
			Sex:     "male",
		}
		//str, _ := json.Marshal(data)
		c.JSON(http.StatusOK, data) //json序列化，通过反射来实现，所以结构体里面的字段首字母要大写
	})

	r.Run()

	//r.Run(":8081")
	/*	//设置返回路由引擎
		r := gin.Default()

		r.GET("/json", func(c *gin.Context) {
			//方法一：map
			// data := map[string]interface{}{
			//     "name":    "小魔仙",
			//     "message": "hello golang",
			//     "age":     13,
			//     "sex":     "男",
			// }
			c.JSON(http.StatusOK, gin.H{
				"name":    "小魔仙",
				"message": "hello golang",
				"age":     13,
				"sex":     "女",
				"b1":      " 今天是2021年08-30",
			})
		})*/

	/*	//结构体返回，灵活使用tag来对结构体字段做定制化操作
		type msg struct {
			Name    string
			Message string
			Age     int
		}
		r.GET("/another_json", func(c *gin.Context) {
			data := msg{
				Name:    "乔四美",
				Message: "七七",
				Age:     77,
			}
			c.JSON(http.StatusOK, data) //json的序列化，通过反射 ，所以结构体内的字段首字母要大写
		})

		r.Run(":9090")*/
}
