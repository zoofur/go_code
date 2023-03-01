package main

import "fmt"

type A interface {
	Get()
}

type B interface {
	A
	Hello()
}

type T struct {

}

func (t T) Get() {
	fmt.Println("类型T")
}

func (t T) Hello() {
	fmt.Println("Hello World")
}

func main() {
	var y A = T{}
	isAssert(y)
}

func isAssert(a A) {
	a.Get()
	b := a.(B)
	b.Hello()
}
