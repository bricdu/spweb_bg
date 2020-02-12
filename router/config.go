package router

import (
	"fmt"
	"io"
	"test2/model"
	"text/template"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/bricdu/utils"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if debug {
		t.templates = template.Must(template.ParseFiles("./views/login.html", "./views/index.html"))
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

var renderer = &TemplateRenderer{
	templates: template.Must(template.ParseFiles("./views/login.html", "./views/index.html")),
}

//ServerHeader 中间件
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		tokenString := ctx.FormValue("token")

		//log.Println(tokenString)
		claims := model.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("123"), nil
		})
		//log.Println(err)
		//log.Println(token.Valid)
		if err == nil && token.Valid {
			//验证通过
			ctx.Set("uid", claims.Id)
			fmt.Println("ss")
			return next(ctx)
		}
		//验证失败
		return ctx.JSON(utils.ErrJwt("验证失败"))

	}
}
