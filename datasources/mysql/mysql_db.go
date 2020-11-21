package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysqlUsersUsername = "mysql_username"
	mysqlUsersPassword = "mysql_password"
	mysqlUsersHost     = "mysql_host"
	mysqlUsersSchema   = "mysql_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	var err error
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		username, password, host, schema,
	)
	Client, err = sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully connected")
}
