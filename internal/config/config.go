package config

import "time"

type Config struct {
	Range    int
	Interval int
	Duration time.Duration
}

func DefaultConfig() Config {
	return Config{
		Range:    3,
		Interval: 50,
		Duration: 10 * time.Second,
	}
}
