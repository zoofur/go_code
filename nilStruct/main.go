package main

import "fmt"

type MyInterface interface {
}

type MyInterfaceImplement struct {
	name string
}

type MyStruct struct {}

// go run -gcflags="-m" main.go
func main() {
	a := new(struct{})
	b := new(struct{})
	println(a, b, a == b) // 空结构体不能相互比较
	println(struct{}{} == struct{}{})

	c := new(struct{})
	d := new(struct{})
	fmt.Println(c, d) // 导致内存逃逸
	println(c, d, c == d)

	g := new(MyStruct)
	h := new(MyStruct)
	println(g, h, g == h)
	println(MyStruct{} == MyStruct{})

	e := new(interface{})
	f := new(interface{})
	println(e, f, e == f) // 接口比较什么?

	i := &MyInterfaceImplement{}
	j := &MyInterfaceImplement{}
	println(i, j, i.name == j.name)
}
