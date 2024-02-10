package config

import "time"

type Config struct {
	SumDelay      time.Duration `json:"sumDelay"`
	DiffDelay     time.Duration `json:"diffDelay"`
	MultiplyDelay time.Duration `json:"multiplyDelay"`
	DevideDelay   time.Duration `json:"devideDelay"`
}

func NewConfig() *Config {
	defaultDelay := 15 * time.Second
	return &Config{
		SumDelay:      defaultDelay,
		DiffDelay:     defaultDelay,
		MultiplyDelay: defaultDelay,
		DevideDelay:   defaultDelay,
	}
}
