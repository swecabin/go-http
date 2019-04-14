package http

import (
	"github.com/FenixAra/go-util/log"
)

type HttpWrapper interface {
}

func New(config *Config, l *log.Logger) HttpWrapper {
	return &httpwrapper{
		c: config,
		l: l,
	}
}
