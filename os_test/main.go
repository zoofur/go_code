package main

import (
	"fmt"
	"os"
)

func main() {
	DevNull, err := os.Create("/dev/null")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	fmt.Println(DevNull)
}
