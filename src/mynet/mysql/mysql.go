package mysql

import (
	"database/sql"
	"log"
	"mynet"

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
	stmt, err := db.Prepare(`INSERT INTO chatlog (sendid,content,time) values (?,?,now())`)
	checkErr(err)
	_, err = stmt.Exec(sendid, content)
	if checkErr(err) {
		return false
	}
	return true
}

func GetUserForEmail(email string) *mynet.User {
	rows, err := db.Query(`select * from user where email = ?`, email)
	checkErr(err)
	if !checkErr(err) {
		for rows.Next() {
			var id int
			var name string
			var passwd string
			var friends string
			var other string
			rows.Columns()
			err = rows.Scan(&id, &name, &passwd, &friends, &other)
			checkErr(err)
			user := &mynet.User{Id: id, Name: name, Passwd: passwd, Friends: friends, Other: other}
			return user
		}
	}
	return nil
}

func checkErr(err error) bool {
	if err != nil {
		log.Println("数据库操作出错")
		log.Panic(err)
		return true
	}
	return false
}
