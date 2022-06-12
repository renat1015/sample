package cmd

import (
	"github.com/fatih/structs"
	"sample/db"
	"sample/server"
)

type config struct {
	Production   bool           `mapstructure:"production"`
	LogLevel     string         `mapstructure:"log_level"`
	ServerConfig *server.Config `mapstructure:"server"`
	DBConfig     *db.Config     `mapstructure:"db"`
}

func (cfg *config) validate() error {
	/*if err := cfg.DBConfig.validate(); err != nil {
		return errors.New("db config: " + err.Error())
	}*/
	return nil
}

func defaultConfig() *config {
	return &config{
		Production:   false,
		LogLevel:     "DEBUG",
		ServerConfig: server.DefaultConfig(),
		DBConfig:     db.DefaultConfig(),
	}
}

func configValues(s interface{}) []string {
	res := make([]string, 0, 1)

	fields := structs.Fields(s)

	for _, field := range fields {
		tag := field.Tag("mapstructure")

		if structs.IsStruct(field.Value()) {
			arr := configValues(field.Value())

			for _, t := range arr {
				res = append(res, tag+"."+t)
			}
		} else {
			res = append(res, tag)
		}
	}

	return res
}
