package main

import (
	"errors"
	"fmt"
)

func getWarpError() {
	err1 := errors.New("原来的错误")
	err2 := fmt.Errorf("fmt.Errorf 进行包装 %w", err1)
	fmt.Println(err2)
}
