package wallet

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"wallet/controllers"
)

func routerSet() {
	e := echo.New()
	// 跨域请求配置
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderAccept, echo.HeaderOrigin, echo.HeaderContentType},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST},
		AllowCredentials: true,	// 允许 cookie
		MaxAge: 43200})) // 预检结果能保留 12h
	g := e.Group("/wallet")
	{
		g.POST("/register", controllers.Register)
	}
}

func main() {
\
}