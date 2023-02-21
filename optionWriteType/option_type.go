package mian

import (
	"errors"
	"time"
)

type Option func(options *Options) error

var defaultPort = 8080

type Options struct {
	Addr string
	Port *int
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
	if addr == "" {
		return nil, errors.New("addr is empty")
	}
	var options Options
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

	options.Addr = addr
	return &options, nil
}

func (o *Options) Run() {

}
