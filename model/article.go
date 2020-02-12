package model

import (
	"errors"
	"time"
)

type Article struct {
	Id        int64     `json:"id,omitempty"`
	Cid       int64     `json:"cid,omitempty"`
	ClassName string    `json:"class_name,omitempty" db:"-"`
	Uid       int64     `json:"uid,omitempty"`
	Title     string    `json:"title,omitempty"`
	Origin    string    `json:"origin,omitempty"`
	Author    string    `json:"author,omitempty"`
	Content   string    `json:"content,omitempty"`
	Hits      int64     `json:"hits,omitempty"`
	Utime     time.Time `json:"utime,omitempty"`
	Ctime     time.Time `json:"ctime,omitempty"`
}

//ArticleCount 文章总数
func ArticleCount() int {
	count := 0
	Db.Get(&count, "select count(*) from article")
	return count
}

//ArticlePage 分页数据
func ArticlePage(pi, ps int) ([]Article, error) {
	mods := make([]Article, 0, ps)
	err := Db.Select(&mods, "select * from article order by id desc limit ?,?", (pi-1)*ps, ps)

	for i := 0; i < len(mods); i++ {

		mods[i].ClassName = ClassNameById(mods[i].Cid)
		mods[i].Content = ""
	}

	return mods, err
}

func iscidin(cid int64, arr []int64) bool {
	for i := 0; i < len(arr); i++ {
		if cid == arr[i] {
			return true
		}
	}
	return false
}

//ArticleDel 删除
func ArticleDel(id int64) error {
	tx, err := Db.Begin()
	if err != nil {
		return err
	}
	result, err := tx.Exec("delete from article where id =?", id)
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

//ArticleAdd 添加文章
func ArticleAdd(mod *Article) error {
	tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	result, err := tx.NamedExec("insert into article(title,author,cid,content,hits,ctime,utime,origin,uid) values (:title,:author,:cid,:content,:hits,:ctime,:utime,:origin,:uid)", mod)
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
