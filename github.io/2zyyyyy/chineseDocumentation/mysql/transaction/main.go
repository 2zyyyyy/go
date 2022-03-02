package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// MySQL 事务

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:zhangyun..@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func insert(username, sex, email string) (sql string) {
	sql = fmt.Sprintf("insert into person(username, sex, email)values(%s, %s, %s)", username, sex, email)
	fmt.Println("sql:", sql)
	return sql
}

func main() {
	conn, err := Db.Begin()
	defer Db.Close()
	if err != nil {
		fmt.Println("begin failed, err", err)
		return
	}
	r, err := conn.Exec(insert("'stu_004'", "'man'", "'stu_004@qq.com'"))
	if err != nil {
		fmt.Println("exec failed, err:", err)
		_ = conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, err:", err)
		_ = conn.Rollback()
		return
	}
	fmt.Println("insert success", id)

	r, err = conn.Exec(insert("'stu_005'", "'man'", "'tu_005@qq.com'"))
	if err != nil {
		fmt.Println("exec failed, err:", err)
		_ = conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, err:", err)
		_ = conn.Rollback()
		return
	}
	fmt.Println("insert success", id)

	// 提交事务
	err = conn.Commit()
	if err != nil {
		fmt.Println("commit failed, err", err)
		return
	}
}
