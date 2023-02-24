package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 4; i++ {
		queryAll()
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
}

