package router

import (
	"test2/control"

	"github.com/labstack/echo"
)

//AdminRouter AdminRouter
func AdminRouter(admin *echo.Group) {
	admin.GET("/index", control.AdminIndexView)
	admin.POST("/class/add", control.ClassAdd)
	admin.POST("/class/edit", control.ClassEdit)
	admin.GET("/class/del/:id", control.ClassDel) //path
	admin.GET("/user/page", control.UserPage)
	admin.GET("/user/del/:id", control.UserDel) //path
	admin.POST("/user/add", control.UserAdd)
	admin.GET("/user/get/:id", control.UserGet)
	admin.POST("/user/edit", control.UserEdit)
	admin.GET("/article/del/:id", control.ArticleDel) //path
	admin.POST("/article/add", control.ArticleAdd)

}
