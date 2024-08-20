package spectator

import "github.com/Netflix/spectator-go/v2/spectator/writer"

// Option is implemented by types that allow extra customization.
type Option interface {
	apply(conf *Config)
}

type configOptionFunc func(conf *Config)

func (c configOptionFunc) apply(conf *Config) {
	if c != nil {
		c(conf)
	}
}

// WithWriter returns an Option that adds a custom writer.
func WithWriter(w writer.Writer) Option {
	return configOptionFunc(func(conf *Config) {
		conf.location = "custom"
		conf.writer = w
	})
}
