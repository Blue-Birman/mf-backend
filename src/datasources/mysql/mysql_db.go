package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB

	host     = os.Getenv("DB_HOST")
	username = "root"
	password = ""
	database = "mf_backend"
)

func init() {
	var err error
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		username, password, host, database,
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
