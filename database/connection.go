package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {

	// GET CONNECTION TO DB
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBNAME"))
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
