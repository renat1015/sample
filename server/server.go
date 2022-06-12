package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"sample/service"
)

type server struct {
	conf    *Config
	service *service.Service
}

func NewServer(conf *Config, service *service.Service) error {
	e := echo.New()
	e.Logger.SetLevel(log.OFF)

	return e.Start(fmt.Sprintf(":%d", conf.Port))
}
