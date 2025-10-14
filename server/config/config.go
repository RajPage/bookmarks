package config

import (
	"bookmarks/server/common"
	"bookmarks/server/config/environment"
	"sync"
)

type Config struct {
	Environment string
	LogLevel    string
	Port        string
}

var (
	singletonConfig *Config
	once            sync.Once
)

func InitializeConfig() *Config {
	once.Do(func() {
		singletonConfig = &Config{
			Environment: common.GetEnvValue(environment.EnvServerEnvironment),
			LogLevel:    common.GetEnvValue(environment.EnvServerLogLevel),
			// Port:        common.GetEnvValue(environment.EnvServerPort),
		}
	})
	return singletonConfig
}

func GetConfig() *Config {
	if singletonConfig == nil {
		panic("Config not initialized. Call InitializeConfig first.")
	}
	return singletonConfig
}
