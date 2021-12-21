package main

import (
	"fmt"
	"log"
	"net/http"
)

//也就是说，对象的实例必须定义为指针的类型，然后才能传递正确的地址，否则ctr参数只是对象的一个副本;若是ctr是结构体类型变量永远只是一个副本传递不了真正的地址
type Counter struct {
	n    int
	name string
}

func (ctr *Counter) ServeHTTP(c http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(c, "%08x\n", ctr)
	ctr.n++
	fmt.Fprintf(c, "张总今年")
	fmt.Fprintf(c, "贵庚 = %d岁了\n", ctr.n)
}

func main() {

	http.Handle("/counter", new(Counter))
	log.Fatal("ListenAndServe: ", http.ListenAndServe(":80", nil))

}
