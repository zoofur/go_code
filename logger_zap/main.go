package main

import (
	"go.uber.org/zap"
	"logger_zap/zlog"
)

type User struct {
	Name string
}

func main() {

	user := &User{
		Name: "Kevin",
	}
	zlog.Info("test log", zap.Any("user", user))
}
