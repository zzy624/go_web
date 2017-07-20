package api

import (
	"mycode/go_web/app/api/user"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	// 兼容老接口，返回结构不变
	g := e.Group("v1")
	g.GET("/name", user.GetName)
}
