package config

import (
	"github.com/charmbracelet/log"
	"github.com/gobuffalo/envy"
)

const defaultLogLevel log.Level = log.ErrorLevel

type Config struct {
	spotify Spotify
	logging Logging
}

type Spotify struct {
	ClientID     string
	ClientSecret string
}

type Logging struct {
	Level log.Level
}

func GetLogConfig() Logging {
	level, err := log.ParseLevel(envy.Get("APP_LOG_LEVEL", ""))
	if err != nil {
		level = defaultLogLevel
	}

	return Logging{
		Level: level,
	}
}

func NewConfig() *Config {
	err := envy.Load()
	if err != nil {
		log.Fatal("failed to load environment variables")
	}
	return &Config{
		logging: GetLogConfig(),
	}
}

func GetLogging() Logging {
	return NewConfig().logging
}
