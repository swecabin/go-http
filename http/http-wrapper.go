package http

type HttpWrapper interface {
}

func New(config *Config) HttpWrapper {
	return &http{}
}
