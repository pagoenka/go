package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// Prepare statement for reading data
	rows, err := db.Query("SELECT * from emp")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Printf("%s \n", name)
	}
}
