package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func init() {
	//dataSource := "root:root@tcp(localhost:3306)/turnos-odontologia"
	dataSource := "root:root@tcp(localhost:3308)/database_back_iii"
	var err error
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	fmt.Println("database configured")
}
