package router

import (
	"test2/control"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// const modeDebug = 1
// const modeProd = 2

var debug = true

//Run 运行
func Run() {
	app := echo.New()
	app.HideBanner = true
	app.Renderer = renderer

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	app.Static("/static", "static")
	app.Static("/views", "views")
	app.GET("/", control.Index)
	app.GET("/login", control.LoginView)

	//adm := app.Group("/admin", ServerHeader)
	//adm.GET("/index", control.AdminIndexView)
	api := app.Group("/api")
	ApiRouter(api)
	admin := app.Group("/admin", ServerHeader)
	AdminRouter(admin)
	app.Start(":8081")
}
