package main

import "fmt"

func main() {
	m := man{
		user: user{
			name: "ylk",
			age:  22,
		},
		sex: "男",
	}
	fmt.Println(getPeopleName(&m))
}
