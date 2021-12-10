package router

import (
	"github.com/gin-gonic/gin"
	"github.com/halhal23/shoezoo/gateway/handler"
	"github.com/halhal23/shoezoo/gateway/middleware"
)

const (
	Ping = "/ping"
)

func NewRouter() *gin.Engine {
	e := gin.Default()
	e.Use(middleware.Cors())
	e.GET(Ping, handler.Ping())
	return e
}
