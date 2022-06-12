package server

import "time"

type Config struct {
	Port           uint          `mapstructure:"port"`
	JWTSecret      string        `mapstructure:"jwt_secret"`
	UserJWTExpires time.Duration `mapstructure:"user_jwt_expires"`
}

func DefaultConfig() *Config {
	return &Config{
		Port:           8080,
		JWTSecret:      "secret",
		UserJWTExpires: 24 * time.Hour,
	}
}
