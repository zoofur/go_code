package main

import (
	"math/rand"
	"time"
)

func queryAll() int {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func() { ch <- query() }()
	}
	return <-ch
}

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}
