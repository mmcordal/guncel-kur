package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	Server ServerConfig
	Redis  RedisConfig
}

type ServerConfig struct {
	Port string
}

type RedisConfig struct {
	Host string
	Port string
}

func setDefaults() {
	viper.SetDefault("server.port", "3000")

	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
}

func Setup() (*Config, error) {
	setDefaults()

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf(".env bulunamadı, env değişkenlerinden okunacak")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	config = &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}

func Get() *Config {
	if config == nil {
		panic("config yüklenmedi")
	}
	return config
}
