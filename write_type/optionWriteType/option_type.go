package mian

import (
	"errors"
	"time"
)
/*
	每个公共字段配置一个函数，实现起来比较优雅。
 */
type Option func(options *Options) error

var defaultPort = 8080

type Options struct {
	Addr string
	Port *int	// Port字段为指针的原因：当Port没有被赋值时，此时缺省值为0，无法判断是否为用户设置的端口。将其字段设置为指针，可以对nil进行判断。
	Timeout time.Duration
}

func WithPort(port int) Option {
	return func(options *Options) error {
		if port < 0 {
			return errors.New("port is err")
		}
		options.Port = &port
		return nil
	}
}

func WithTimeout(time time.Duration) Option {
	return func(options *Options) error {
		options.Timeout = time
		return nil
	}
}

func NewOptions(addr string, opts ...Option) (*Options, error) {
	options := Options{
		Addr: addr,
	}
	if addr == "" {
		return nil, errors.New("addr is empty")
	}
	var err error
	for _, opt := range opts {
		err = opt(&options)
		if err != nil {
			return nil, err
		}
	}

	if options.Port == nil {
		options.Port = &defaultPort
	}
	return &options, nil
}

func (o *Options) Run() {

}
