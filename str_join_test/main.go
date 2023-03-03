package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Builder2String([]string{
		"拼接",
		"测试",
		"函数",
	}))
}

func Builder2String(str [] string) string {
	b := strings.Builder{}

	n := 0
	for i := 0; i < len(str); i++ {
		n += len(str[i])
	}

	b.Grow(n)
	for i := 0; i < len(str); i++ {
		b.WriteString(str[i])
	}
	return b.String()
}
