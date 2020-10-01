package config

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//konstanta
const (
	username string = "root"
	password string = ""
	database string = "siswa"
)
var (
	setting = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

//koneksi
func Mysql() (*sql.DB, error){
	db, err := sql.Open("mysql", setting)

	if err != nil {
		return nil, err
	}
	return db, nil
}