package standard

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var db struct {
	Dns string
}

func TestMain(m *testing.M) {
	fmt.Println("enter TestMain function")
	setup()
	db.Dns = os.Getenv("DATABASE_DNS")
	if db.Dns == "" {
		db.Dns = "root:123456@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local"
	}

	flag.Parse()
	exitCode := m.Run()

	db.Dns = ""

	teardown()
	// 退出
	os.Exit(exitCode)
}

func TestDatabase1(t *testing.T) {
	fmt.Println("enter TestDatabase1 function:", db.Dns)
}

func TestDatabase2(t *testing.T) {
	fmt.Println("enter TestDatabase2 function:", db.Dns)
}
