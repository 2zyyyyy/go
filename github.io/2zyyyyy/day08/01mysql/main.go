package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
)

// golang 连接数据库示例

// 定义全局对象db
var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

// 初始化数据库
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)

	// 设置数据库连接池的最大空闲线程数
	db.SetMaxOpenConns(5)
	return nil
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Print("init db success!!\n")

	queryRow(2)
	queryMore(0)
	deleteRow()
	insertRow()
	updateRow()
}

// 查询单条数据
func queryRow(id int) {
	// 1.单条查询
	sqlStr := "select id, name, age from user where id =?;"
	// 2.执行（演示超过最大连接数）
	// for i := 0; i < 20; i++ {
	// 	fmt.Printf("第%d次查询\n", i)
	// 	db.QueryRow(sqlStr, 1)
	// }
	// 3.拿到结果
	var u user
	// 非常重要。确保QuertRow之后调用Scan方法，否则持有的数据库连接不会被释放
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 查询多行数据
func queryMore(id int) {
	sqlStr := "select id, name, age from user where id >?;"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("sacn failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 插入数据
func insertRow() {
	sqlStr := "INSERT INTO `sql_test`.`user` (`id`, `name`, `age`) VALUES ('1', '插入的数据', '12580');"
	res, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get insert id failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d\n", id)

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get insert iRowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("Affected:%d\n", affected)
}

// 更新数据
func updateRow() {
	sqlStr := "UPDATE `sql_test`.`user` SET age=? WHERE ID = ?;"
	res, err := db.Exec(sqlStr, 9000, 6)
	if err != nil {
		fmt.Printf("UPDATE failed,err:%v\n", err)
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get UPDATE iRowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("Affected:%d\n", affected)
}

// 删除数据
func deleteRow() {
	sqlStr := "DELETE FROM `sql_test`.`user` WHERE ID = ?;"
	res, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("UPDATE failed,err:%v\n", err)
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get UPDATE iRowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("Affected:%d\n", affected)
}
