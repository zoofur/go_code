package main

import (
	"errors"
	"time"
)

/*
	需创建一个 ConfigBuilder 作为配置项传入新建服务中，即使没有需要配置的项也需要创建一个空结构体传入。
	对 error 的处理延迟到了 ConfigBuilder 的 Builder 方法中进行，无法做到对字段粒度的判断。
	相对于 Option 方式实现的优点，其方法执行是有顺序的。
 */

var defaultPort = 8080

type Config struct {
	Addr string
	Port *int
	Timeout time.Duration
}

type ConfigBuilder struct {
	Port *int
	Timeout time.Duration
}

func (cb *ConfigBuilder) BuildWithPort(port int) *ConfigBuilder {
	cb.Port = &port
	return	cb
}

func (cb *ConfigBuilder) BuildWithTimeout(timeout time.Duration) *ConfigBuilder {
	cb.Timeout = timeout
	return cb
}

func (cb *ConfigBuilder) Build() *Config {
	if cb.Port == nil {
		cb.Port = &defaultPort
	}

	c := Config{
		Port: cb.Port,
		Timeout: cb.Timeout,
	}
	return &c
}

func (c *Config) Run() {

}

func NewServer(config *Config, addr string) (*Config, error) {
	if addr == "" {
		return nil, errors.New("addr is empty")
	}

	config.Addr = addr
	return config, nil
}
