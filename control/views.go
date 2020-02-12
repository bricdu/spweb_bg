package control

import "github.com/labstack/echo"

//LoginView LoginView
func LoginView(ctx echo.Context) error {
	return ctx.Render((200), "login.html", nil)
}

//AdminIndexView AdminIndexView
func AdminIndexView(ctx echo.Context) error {
	return ctx.Render((200), "index.html", nil)
}
