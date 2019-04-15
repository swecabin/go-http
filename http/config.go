package http

type Config struct {
	// HTTP timeout in seconds. Default - 10 second
	timeout int
	// No of Retries in count - Default - 3
	retries int
	// Headers for the http request
	headers     map[string]string
	retryFactor float64
}

// Creates and initialises config to default values
func NewConfig() *Config {
	return &Config{
		timeout:     10,
		retries:     3,
		headers:     make(map[string]string),
		retryFactor: 2,
	}
}

// Function to add new HTTP header for all requests. k - key of header (Authorisation etc.)
// v - Value of the header
func (c *Config) AddHeader(k, v string) {
	c.headers[k] = v
}

// Function to Set timeout for each HTTP requests
func (c *Config) SetTimeout(timeout int) {
	c.timeout = timeout
}

// Function to set number of reties
func (c *Config) SetRetries(retries int) {
	c.retries = retries
}

// Function to set retry factor for exponential backoff
func (c *Config) SetRetryFactor(factor float64) {
	c.retryFactor = factor
}
