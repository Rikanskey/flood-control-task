package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	defaultAddr        = "localhost:6379"
	defaultPassword    = ""
	defaultDB          = 0
	defaultRefreshTime = time.Duration(5) * time.Second
	defaultMaxRequests = 100
)

type Config struct {
	App   AppConfig
	Redis RedisConfig
}

type AppConfig struct {
	MaxRequests int
}

type RedisConfig struct {
	Addr        string
	Password    string
	DB          int
	RefreshTime time.Duration
}

func (rc *RedisConfig) withDefaults() (r RedisConfig) {
	if rc != nil {
		r = *rc
	}

	r.Password = defaultPassword
	r.Addr = defaultAddr
	r.DB = defaultDB
	if r.RefreshTime == 0 {
		r.RefreshTime = defaultRefreshTime
	} else {
		r.RefreshTime *= time.Second
	}

	return
}

func (ac *AppConfig) withDefaults() (a AppConfig) {
	if ac != nil {
		a = *ac
	}

	a.MaxRequests = defaultMaxRequests

	return
}

func New(configDir string) (*Config, error) {
	if err := parseConfigFile(configDir); err != nil {
		return nil, err
	}

	var config Config

	err := unmarshall(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}

func parseConfigFile(configDir string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("application")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func unmarshall(config *Config) error {
	if err := viper.UnmarshalKey("db", &config.Redis); err != nil {
		config.Redis.withDefaults()
	}
	if err := viper.UnmarshalKey("app", &config.App); err != nil {
		config.App.withDefaults()
	}

	return nil
}
