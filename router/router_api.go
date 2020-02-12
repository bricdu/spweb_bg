package router

import (
	"test2/control"

	"github.com/labstack/echo"
)

//ApiRouter ApiRouter
func ApiRouter(api *echo.Group) {
	api.POST("/login", control.UserLogin)
	api.GET("/class/all", control.ClassAll)
	api.GET("/class/page", control.ClassPage)
	api.GET("/class/get/:id", control.ClassGet)
	api.GET("/article/page", control.ArticlePage)
}
