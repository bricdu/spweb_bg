package model

import "errors"

//Class 结构体
type Class struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`
}

//ClassPage 分页
func ClassPage(pi, ps int) ([]Class, error) {
	mods := make([]Class, 0, ps)
	err := Db.Select(&mods, "select * from class limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

//ClassCount 数据总数
func ClassCount() int {
	count := 0
	Db.Get(&count, "select count(id) as count from class")
	return count
}

//ClassAll 查询所有
func ClassAll() ([]Class, error) {
	mods := make([]Class, 0, 8)
	err := Db.Select(&mods, "select * from class ")
	return mods, err
}

//ClassGet 查询单个
func ClassGet(id int64) (*Class, error) {
	mod := &Class{}
	err := Db.Get(mod, "select * from class where id =? limit 1", id)
	return mod, err
}

//ClassAdd 添加
func ClassAdd(mod *Class) error {
	tx, err := Db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("insert into class(`name`,`desc`) value (?,?)", mod.Name, mod.Desc)
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

//ClassEdit 修改
func ClassEdit(mod *Class) error {
	tx, err := Db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("update class set `name`=?,`desc`=? where id=?", mod.Name, mod.Desc, mod.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		tx.Rollback()
		return errors.New("edit Rows Affected < 1")
	}
	tx.Commit()
	return nil
}

//ClassDel 删除
func ClassDel(id int64) error {
	tx, err := Db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("delete from class where id =?", id)
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

//ClassNameById 获取class名称
func ClassNameById(id int64) string {
	name := ""
	Db.Get(&name, "select name from class where id=?", id)
	return name
}
