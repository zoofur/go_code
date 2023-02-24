package main

import "fmt"

type user struct {}

func (u *user) Hello(name string) {
	fmt.Println("你好" + name)
	//u.NoHello()
}

/**
func Hello(u *User, name string) {...}
 */

func (u user) NoHello() {
	fmt.Println("你不好")
}


func main() {
	var u user
	u.Hello("于立凯")
	u.NoHello()
}