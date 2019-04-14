package http

type Config struct {
	// HTTP timeout in seconds. Default - 10 second
	Timeout int
	// No of Retries in count - Default - 3
	Retries int
	// Headers for the http request
	Headers map[string]string
}

func NewConfig() *Config {
	return &Config{
		Timeout: 10,
		Retries: 3,
		Headers: make(map[string]string),
	}
}

func (c *Config) AddHeader(k, v string) {
	c.Headers[k] = v
}
