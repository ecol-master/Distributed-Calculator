package config

import (
	"fmt"
	"time"
)

const (
	WorkerAddress  = "localhost:8000"
	StorageAddress = "localhost:5000"
)

type Config struct {
	SumDelay      time.Duration `json:"sumDelay"`
	DiffDelay     time.Duration `json:"diffDelay"`
	MultiplyDelay time.Duration `json:"multiplyDelay"`
	DevideDelay   time.Duration `json:"devideDelay"`
}

func NewConfig() *Config {
	defaultDelay := 1 * time.Second
	return &Config{
		SumDelay:      defaultDelay,
		DiffDelay:     defaultDelay,
		MultiplyDelay: defaultDelay,
		DevideDelay:   defaultDelay,
	}
}

func (c *Config) CopySettings(newCfg Config) {
	if !newCfg.isValid() {
		return
	}

	c.SumDelay = newCfg.SumDelay
	c.DiffDelay = newCfg.DiffDelay
	c.MultiplyDelay = newCfg.MultiplyDelay
	c.DevideDelay = newCfg.DevideDelay
}

func (c *Config) isValid() bool {
	return c.SumDelay >= 0 && c.DiffDelay >= 0 && c.MultiplyDelay >= 0 && c.DevideDelay > 0
}

func (c *Config) AsString() string {
	configAsString := fmt.Sprintf("{sum: %.2f sec, diff: %.2f sec, multiply: %.2f sec, devide: %.2f sec}", c.SumDelay.Seconds(), c.DiffDelay.Seconds(), c.MultiplyDelay.Seconds(), c.DevideDelay.Seconds())
	return configAsString
}
