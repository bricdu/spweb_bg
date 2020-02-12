package control

import "github.com/labstack/echo"

//Index 首页
func Index(ctx echo.Context) error {
	return ctx.Redirect(302, "/login ")
}
