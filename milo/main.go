package main

import (
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/server/xecho"
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/labstack/echo"
)

type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(eng.ServeHTTP); err != nil {
		xlog.Panic("startup", xlog.Any("err", err))
	}
	return eng
}

func (eng *Engine) ServeHTTP() error {
	server := xecho.StdConfig("http").Build()
	server.GET("/hello", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello World!")
	})
	return eng.Serve(server)
}

func main() {
	eng := NewEngine()
	if err := eng.Run(); err != nil {
		xlog.Error(err.Error())
	}
}
