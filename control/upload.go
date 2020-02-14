package control

import (
	"io"
	"log"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/bricdu/utils"
	"github.com/labstack/echo"
)

//ApiUpload 上传api
func ApiUpload(ctx echo.Context) error {
	//ctx.ParseMultipartForm(1 << 20)
	f, err := ctx.FormFile("upfile")
	if err != nil {
		return ctx.JSON(utils.ErrIpt("上传失败", err.Error()))
	}
	log.Println(f, err)
	src, err := f.Open()
	if err != nil {
		return ctx.JSON(utils.ErrIpt("上传失败", err.Error()))
	}
	defer src.Close()
	os.MkdirAll("static/upload", 0666)
	ext := path.Ext(f.Filename)
	name := "static/upload/" + utils.RandStr(16) + ext
	dst, err := os.Create(name)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("上传失败", err.Error()))
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("上传失败", err.Error()))
	}
	//f.Close()
	//dst.Close()
	//Succ(w, "succ", "/"+name)
	//w.Header().Set("Content-Type", "application/json")
	//w.Write([]byte("{\"original\":\"" + h.Filename + "\",\"state\":\"SUCCESS\",\"title\":\"" + h.Filename + "\",\"url\":\"" + ("/" + name) + "\"}"))
	mod := EditorReply{
		Original: f.Filename,
		State:    "SUCCESS",
		Title:    f.Filename,
		Url:      "/" + name,
	}
	//w.Write(mod.Json())

	return ctx.JSON(utils.Succ("上传成功", mod))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//EditorReply 返回结构体
type EditorReply struct {
	Original string `json:"original,omitempty"`
	State    string `json:"state,omitempty"`
	Title    string `json:"title,omitempty"`
	Url      string `json:"url,omitempty"`
}

// //Json 输出方法
// func (er *EditorReply) Json() []byte {
// 	buf, _ := json.Marshal(er)
// 	return buf
// }
