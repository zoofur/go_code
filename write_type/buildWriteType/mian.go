package main

import (
	"fmt"
	"time"
)

func main() {
	configBuilder := ConfigBuilder{}
	server, err := NewServer(configBuilder.BuildWithPort(9000).BuildWithTimeout(1 * time.Second).Build(), "localhost")
	if err != nil {
		fmt.Println("server fail to run")
	}
	server.Run()
}
