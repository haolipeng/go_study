package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

const (
	userName = "root"
	password = "Happy123!"
	ip       = "192.168.57.138"
	port     = "3306"
	dbName   = "fileserver"
)

var db *sql.DB

func init() {
	//构建连接 "用户名:密码@tcp(ip:端口)/数据库名?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	var err error
	db, err = sql.Open("mysql", path)
	if err != nil {
		fmt.Printf("open mysql error :%s\n", err.Error())
		os.Exit(1)
	}

	db.SetMaxOpenConns(100)
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql err:", err.Error())
		os.Exit(1)
	}
}

//DBConn：返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
