package http

import (
	"github.com/FenixAra/go-util/log"
)

// HttpWrapper is the abstracted interface of http
type HttpWrapper interface {
	MakeRequest(method, url, name string, req, res interface{}) error
}

// New is used to get new HTTP wrapper object
func New(config *Config, l *log.Logger) HttpWrapper {
	return &httpwrapper{
		c: config,
		l: l,
	}
}
