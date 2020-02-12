package model

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id    int    `json:"id,omitempty"`
	Num   string `json:"num,omitempty"`
	Name  string `json:"name,omitempty"`
	Pass  string `json:"pass,omitempty"`
	Phone int    `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	//Status int       `json:"status,omitempty"`
	Ctime time.Time `json:"ctime,omitempty"`
}

//UserLogin 登录
func UserLogin(num string) (User, error) {
	mod := User{}
	err := Db.Get(&mod, "select * from user where num=? limit 1", num)
	return mod, err
}

//UserClaims token 数据
type UserClaims struct {
	Id   int    `json:"id,omitempty"`
	Num  string `json:"num,omitempty"`
	Name string `json:"name,omitempty"`
	jwt.StandardClaims
}

//UserPage 分页数据
func UserPage(pi, ps int) ([]User, error) {
	mods := make([]User, 0, ps)
	err := Db.Select(&mods, "select * from user limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

//UserCount 总数
func UserCount() int {
	count := 0
	Db.Get(&count, "select count(id) as count from user")
	return count
}

//UserDel 删除用户
func UserDel(id int64) error {
	tx, err := Db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("delete from user where id =?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("Rows Affected < 1")
	}
	tx.Commit()
	return nil
}

//UserAdd 添加用户
func UserAdd(mod *User) error {
	tx, err := Db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("insert into user(num,`name`,pass,phone,email,ctime) value (?,?,?,?,?,?)", mod.Num, mod.Name, mod.Pass, mod.Phone, mod.Email, mod.Ctime)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("add Rows Affected < 1")
	}
	tx.Commit()
	return nil
}

//UserExists 用户是否存在
func UserExists(num string) bool {
	mod := User{}
	err := Db.Get(&mod, "select * from user where num = ?", mod.Num)
	if err != nil {
		return false
	}
	return true
}

//UserGet id查用户
func UserGet(id int64) (*User, error) {
	mod := &User{}
	err := Db.Get(mod, "select * from user where id =? limit 1", id)
	return mod, err
}

//UserEdit 修改用户
func UserEdit(mod *User) error {
	tx, err := Db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("update user set `name`=?,pass=?,phone=?,email=? where id = ?", mod.Name, mod.Pass, mod.Phone, mod.Email, mod.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("add Rows Affected < 1")
	}
	tx.Commit()
	return nil
}
