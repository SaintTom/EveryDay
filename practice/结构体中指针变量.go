package main

import (
	"fmt"
	"log"
	"net/http"
)

type Mjperson struct {
	name  string
	age   int
	hobby string
	Friend
}
type ip struct {
	add string
}

type Friend struct {
	name     string
	friendip *ip
}

func (ctr *Mjperson) ServeHTTP(c http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(c, "%08x\n", ctr)
	//ctr.n++
	//fmt.Fprintf(c,"张总今年")
	//fmt.Fprintf(c, "贵庚 = %d岁了\n", ctr.n)
	ctr = new(Mjperson)
	ctr.friendip = new(ip)
	ctr.name = "张总"
	ctr.hobby = "摸鱼"
	ctr.age = 90
	ctr.Friend.name = "凤姐"
	ctr.Friend.friendip.add = "127.0.0.1"
	fmt.Fprintf(c, ctr.name)
	fmt.Fprintf(c, "武功天下第一\n的IPadd")
	fmt.Fprintf(c, ctr.friendip.add)
}

func main() {
	http.Handle("/varpoint", new(Mjperson))
	log.Fatal("ListenAndServe: ", http.ListenAndServe(":81", nil))
}
