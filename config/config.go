package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = ""
	dbname   = "db_golang"
)

var dsn = fmt.Sprintf("%v:%v@/%v", username, password, dbname)

// Connect to MySQL
func ConnectToMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
