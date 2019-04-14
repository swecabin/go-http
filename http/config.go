package http

type Config struct {
	// HTTP timeout in seconds. Default - 10 second
	Timeout int
	// No of Retries in count - Default - 3
	Retries int
	// Headers for the http request
	Headers     map[string]string
	retryFactor float64
}

// Creates and initialises config to default values
func NewConfig() *Config {
	return &Config{
		Timeout:     10,
		Retries:     3,
		Headers:     make(map[string]string),
		retryFactor: 2,
	}
}

// Function to add new HTTP header for all requests. k - key of header (Authorisation etc.)
// v - Value of the header
func (c *Config) AddHeader(k, v string) {
	c.Headers[k] = v
}
