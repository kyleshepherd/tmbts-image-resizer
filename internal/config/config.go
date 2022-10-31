package config

import (
	"github.com/rs/zerolog"
	"go.soon.build/kit/config"
)

// application name
const AppName = "tmbts-image-resizer"

// Config stores configuration options set by configuration file or env vars
type Config struct {
	Log Log
}

// Log contains logging configuration
type Log struct {
	Console bool
	Verbose bool
	Level   string
}

// Default is a default configuration setup with sane defaults
var Default = Config{
	Log{
		Level: zerolog.InfoLevel.String(),
	},
}

// New constructs a new Config instance
func New(opts ...config.Option) (Config, error) {
	c := Default
	v := config.ViperWithDefaults("tmbts-image-resizer")
	err := config.ReadInConfig(v, &c, opts...)
	if err != nil {
		return c, err
	}
	return c, nil
}
