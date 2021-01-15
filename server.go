package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"wallet/controllers"

	"github.com/labstack/echo/middleware"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := echo.New()
	// 跨域请求配置

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderAccept, echo.HeaderOrigin, echo.HeaderContentType},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST},
		AllowCredentials: true,	// 允许 cookie
		MaxAge: 43200})) // 预检结果能保留 12h
	// 一组路由
	g := e.Group("/wallet")
	{
		g.POST("/register", controllers.Register)
		g.File("/html", "./html/main.html")
	}
	// 网页的静态文件
	// 启动服务，平滑关闭
	go func() {
		if err := e.Start(":1998"); err != nil{
			e.Logger.Fatal("Fail to star with error:%v", err)
		}
	}()
	fmt.Println("服务启动成功")
	// 监听停止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	// 留 5s 处理已经接受的请求，然后关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil{
		e.Logger.Fatal("Fail to shutdown with error", err)
	}
}