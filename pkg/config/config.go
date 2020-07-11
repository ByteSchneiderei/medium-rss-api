package config

import "github.com/caarlos0/env"

// EnvVar represents environment variables
type EnvVar struct {
	GoEnv   string `env:"GO_ENV" envDefault:"development"`
	BaseURL string `env:"BASE_URL" envDefault:"localhost"`
	Port    int    `env:"PORT" envDefault:"8080"`

	MediumRSSFeedURL string `env:"MEDIUM_RSSFEED_URL" envDefault:"https://medium.com/feed"`
	MediumProfile    string `env:"MEDIUM_PROFILE,required"`
}

// Parse parses config value from environment variable
func (e *EnvVar) Parse() error {
	return env.Parse(e)
}
