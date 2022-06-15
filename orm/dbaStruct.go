package orm

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DataBaseObject struct {
	Username string
	Password string
	DB       sql.DB
}

func CreateDatabaseObject() {
	var dao = DataBaseObject{
		Username: "root",
		Password: "hb80M!ZYHz",
	}
	s, err := dao.SqlLogin("127.0.0.1", "LOCALTEST")
	if err != nil {
		fmt.Println("sql error", err)
	}
	fmt.Println(s)

	db, err := sql.Open("mysql", s)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// // SQL Queries resulting in multiple Row
	// var (
	// 	id int
	// 	name string
	// )
	// rows, err := db.Query("SELECT id, name FROM test1 where id >= ?", 1)

	// SQL Queries resulting in single Row
}
