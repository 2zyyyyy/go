package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

var db *sql.DB

// 初始化数据库
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置数据库连接池的最大空闲线程数
	db.SetMaxOpenConns(5)
	return
}

func transaction() {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("transaction Begin failed, err:%v\n", err)
		return
	}
	sqlStr1 := "UPDATE user SET age=age-2 where id=9;"
	sqlStr2 := "Update user set age=age+2 where id=8;"

	// 执行SQL1
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		// rollback
		tx.Rollback()
		fmt.Printf("执行sql1出错，需要回滚，错误信息：%v\n", err)
		return
	}
	// 执行SQL2
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		// rollback
		tx.Rollback()
		fmt.Printf("执行sql2出错，需要回滚，错误信息：%v\n", err)
		return
	}
	// 如果SQL1和SQL2都执行成功，就提交本次事务
	err = tx.Commit()
	if err != nil {
		// rollback
		tx.Rollback()
		fmt.Printf("执行tx.Commit()出错，需要回滚，错误信息：%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}

func main() {
	initDB()
	transaction()
}
