package router

import (
	"github.com/labstack/echo"
	"github.com/rs/xid"
)

var (
	guidGenerator = xid.New()
)

func SetupRouters(e *echo.Echo) {
	e.GET("/post", GetPost)
	e.POST("/post", PostPost)
	e.PUT("/post", PutPost)
	e.DELETE("/post", DeletePost)
}
