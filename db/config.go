package db

type Config struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SslMode  string `mapstructure:"sslmode"`
}

func DefaultConfig() *Config {
	return &Config{}
}
