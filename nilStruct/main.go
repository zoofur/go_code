package main

import "fmt"

type MyInterface interface {
}

type MyInterfaceImplement struct {
	name string
}

func (m *MyInterfaceImplement) Get() string {
	return ""
}

type MyStruct struct {}

// go run -gcflags="-m" main.go
func main() {
	a := new(struct{})
	b := new(struct{})
	println(a, b, a == b) // 空结构体不能相互比较

	c := new(struct{})
	d := new(struct{})
	fmt.Println(c, d) // 导致内存逃逸
	println(c, d, c == d)

	g := new(MyStruct)
	h := new(MyStruct)
	println(g, h, g == h)

	e := new(MyInterface)
	f := new(MyInterface)
	println(e, f, e == f) // 指针比较什么? 类型和值?

	i := &MyInterfaceImplement{}
	j := &MyInterfaceImplement{}
	println(i, j, i.name == j.name)
}
