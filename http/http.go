package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"time"

	"github.com/FenixAra/go-util/log"
)

type httpwrapper struct {
	c *Config
	l *log.Logger
}

// Function to make HTTP request. method - HTTP method like GET, POST.
// url - HTTP Request URL. req - Request of HTTP request
// res - Pointer to response object
func (h *httpwrapper) MakeRequest(method, url string, req, res interface{}) error {
	client := &http.Client{
		Timeout: time.Duration(h.c.Timeout) * time.Second,
	}

	var retries int
	for {
		var body []byte
		var err error
		if req == nil {
			body, err = json.Marshal(req)
			if err != nil {
				h.l.Errorf("Unable to marshal req: %+v. Err: %+v", req, err)
				return err
			}
		}

		request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			h.l.Errorf("Unable to create new HTTP Req. Err: %+v", err)
			return err
		}

		for k, v := range h.c.Headers {
			request.Header.Set(k, v)
		}

		response, err := client.Do(request)
		if err != nil {
			h.l.Errorf("Unable to send HTTP Req. Err: %+v", err)
			time.Sleep(time.Second * time.Duration(int(math.Pow(h.c.retryFactor, float64(retries)))))
			retries++
			if retries > h.c.Retries {
				continue
			}

			return err
		}

		if response.StatusCode >= http.StatusInternalServerError {
			h.l.Errorf("Response code is greater than 500. Code: %d", response.StatusCode)
			time.Sleep(time.Second * time.Duration(int(math.Pow(h.c.retryFactor, float64(retries)))))
			retries++
			if retries > h.c.Retries {
				continue
			}

			return err
		}

		if response.StatusCode >= http.StatusBadRequest {
			h.l.Errorf("Response code is between 400 To 499. Code: %d", response.StatusCode)
			return err
		}

		if res != nil {
			content, err := ioutil.ReadAll(response.Body)
			if err != nil {
				h.l.Errorf("Unable to read HTTP Response. Err: %+v", err)
				return err
			}

			err = json.Unmarshal(content, &res)
			if err != nil {
				h.l.Errorf("Unable to unmarshal HTTP Response. Err: %+v", err)
				return err
			}
		}

		response.Body.Close()
		return nil
	}
}
