package main

import (
	"net/http"

	http_wrap "github.com/FenixAra/go-http/http"
	"github.com/FenixAra/go-util/log"
)

func main() {
	config := log.NewConfig("TestApp")
	l := log.New(config)

	cfg := http_wrap.NewConfig()
	cfg.SetRetries(10)
	cfg.SetTimeout(5)
	wrapper := http_wrap.New(cfg, l)

	wrapper.MakeRequest(http.MethodGet, "https://www.google.com", "Google", nil, nil)
}
