package mian

import (
	"fmt"
	"time"
)

func main() {
	server, err := NewOptions("localhost",
		WithPort(9000),
		WithTimeout(1 * time.Second))
	if err != nil {
		fmt.Println(err)
	}
	server.Run()
}
