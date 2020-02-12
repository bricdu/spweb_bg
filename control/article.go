package control

import (
	"strconv"
	"test2/model"
	"time"

	"github.com/bricdu/utils"
	"github.com/labstack/echo"
)

//ArticlePage 文章分页
func ArticlePage(ctx echo.Context) error {
	ipt := Page{}
	err := ctx.Bind(&ipt)
	if err != nil {
		ctx.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if ipt.Ps < 1 || ipt.Ps > 50 {
		ipt.Ps = 10
	}
	if ipt.Pi < 1 {
		ipt.Pi = 1
		//return ctx.JSON(utils.ErrIpt("输入数据有误", "pi <1"))
	}
	count := model.ArticleCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据"))
	}
	mods, err := model.ArticlePage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Page("新闻数据", mods, count))
}

//ArticleDel 文章删除
func ArticleDel(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据错误", err.Error()))
	}
	err = model.ArticleDel(id)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("删除成功"))
}

//ArticleAdd 添加文章
func ArticleAdd(ctx echo.Context) error {
	ipt := model.Article{}
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入数据错误", err.Error()))
	}
	if ipt.Title == "" {
		return ctx.JSON(utils.ErrIpt("标题不能为空"))
	}

	ipt.Ctime = time.Now()
	ipt.Utime = ipt.Ctime
	ipt.Uid, _ = ctx.Get("uid").(int64)
	//ipt.Status = 1
	//如果存在然后报错

	err = model.ArticleAdd(&ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("添加成功"))
}
