package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "fudali:fudali133@tcp(139.129.4.187:3306)/godoob?charset=utf8")
	checkErr(err)
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func InsertChatContent(sendid string, content string) bool {
	stmt, err := db.Prepare(`INSERT INTO chatlog (sendid,content) values (?,?)`)
	checkErr(err)
	_, err = stmt.Exec(sendid, content)
	if checkErr(err) {
		return false
	}
	return true
}

func checkErr(err error) bool {
	if err != nil {
		log.Println("数据库操作出错")
		log.Panic(err)
		return true
	}
	return false
}
