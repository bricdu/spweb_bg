package model

import (
	"bufio"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	fmt.Println("数据库用户名：")
	dbuser := bufio.NewScanner(os.Stdin)
	dbuser.Scan()
	fmt.Println("数据库密码：")
	dbpass := bufio.NewScanner(os.Stdin)
	dbpass.Scan()
	con := dbuser.Text() + ":" + dbpass.Text() + "@tcp(127.0.0.1:3306)/news?charset=utf8&parseTime=true"
	db, err := sqlx.Open(`mysql`, con)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalln(err.Error())
	}
	Db = db
}
