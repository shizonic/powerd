package main

import (
	"time"
)

type Config struct {
	checkInterval time.Duration
	warnOnLevel   int
	warnMessage   string
	execOnLevel   int
	execCommand   string
}

func NewConfig() *Config {
	config := &Config{
		checkInterval: 60,
		warnOnLevel:   10,
		warnMessage:   "Battery level at 10%",
		execOnLevel:   5,
		execCommand:   "poweroff",
	}
	return config
}
