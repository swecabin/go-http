# go-http
HTTP Wrapper in golang along with mock

To get the latest package: 

```sh
go get -u github.com/FenixAra/go-http/http
```

## Usage
```
package main

import (
	"net/http"

	http_wrap "github.com/FenixAra/go-http/http"
	"github.com/FenixAra/go-util/log"
)

func main() {
	config := log.NewConfig("", "Debug", "Full", "TestApp", "", "", "")
	l := log.New(config)

	cfg := http_wrap.NewConfig()
	cfg.SetRetries(10)
	cfg.SetTimeout(5)
	wrapper := http_wrap.New(cfg, l)

	wrapper.MakeRequest(http.MethodGet, "https://www.google.com", nil, nil)
}
```