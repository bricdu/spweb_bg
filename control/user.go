package control

import (
	"fmt"
	"strconv"
	"test2/model"
	"time"

	"github.com/bricdu/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type login struct {
	Num  string `json:"num,omitempty"`
	Pass string `json:"pass,omitempty"`
}

//UserLogin 登录
func UserLogin(ctx echo.Context) error {
	ipt := login{}
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))

	}
	mod, err := model.UserLogin(ipt.Num)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("用户名或密码错误"))
	}
	if mod.Pass != ipt.Pass {
		return ctx.JSON(utils.ErrIpt("用户名或密码错误"))
	}
	// Create the Claims
	claims := model.UserClaims{
		Id:   mod.Id,
		Num:  mod.Num,
		Name: mod.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("123"))
	fmt.Printf("%v %v", ss, err)
	return ctx.JSON(utils.Succ("登陆成功", ss))
}

//Page 分页结构体 pi 页码 ps 页容量
type Page struct {
	Pi int `json:"pi,omitempty"`
	Ps int `json:"ps,omitempty"`
}

//UserPage UserPage
func UserPage(ctx echo.Context) error {
	ipt := Page{}
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	if ipt.Ps < 1 || ipt.Ps > 50 {
		ipt.Ps = 6
	}
	if ipt.Pi < 1 {
		ipt.Pi = 1
		//return ctx.JSON(utils.ErrIpt("输入数据有误", "pi <1"))
	}
	count := model.UserCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	mods, err := model.UserPage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Page("用户数据", mods, count))
}

//UserDel 删除用户
func UserDel(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据错误", err.Error()))
	}
	uid, _ := ctx.Get("uid").(int64)
	if uid == id {
		return ctx.JSON(utils.Fail("不能删除自己"))
	}
	err = model.UserDel(id)
	if err != nil {
		return ctx.JSON(utils.Fail("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("删除成功"))
}

//UserAdd 添加用户
func UserAdd(ctx echo.Context) error {
	ipt := &model.User{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据错误", err.Error()))
	}
	if ipt.Num == "" {
		return ctx.JSON(utils.ErrIpt("用户账号不能为空"))
	}
	if ipt.Name == "" {
		return ctx.JSON(utils.ErrIpt("用户名不能为空"))
	}
	if ipt.Pass == "" {
		return ctx.JSON(utils.ErrIpt("密码不能为空"))
	}
	ipt.Ctime = time.Now()
	//ipt.Status = 1
	//如果存在然后报错
	if model.UserExists(ipt.Num) {
		return ctx.JSON(utils.ErrIpt("当前账号已经存在"))
	}
	err = model.UserAdd(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("添加成功"))
}

//UserGet id查用户
func UserGet(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据错误", err.Error()))
	}
	mod, err := model.UserGet(id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("查询失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("查询成功", mod))
}

//UserEdit 用户信息修改
func UserEdit(ctx echo.Context) error {
	ipt := &model.User{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据错误", err.Error()))
	}

	if ipt.Name == "" {
		return ctx.JSON(utils.ErrIpt("用户名不能为空"))
	}
	if ipt.Pass == "" {
		return ctx.JSON(utils.ErrIpt("密码不能为空"))
	}

	//ipt.Status = 1
	//如果存在然后报错

	err = model.UserEdit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("添加成功"))
}
