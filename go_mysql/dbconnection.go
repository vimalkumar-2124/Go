package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type OutputDb struct {
	name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:Abcdefghij2124$@tcp(127.0.0.1:3306)/students")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// insert, err := db.Query(
	// 	"INSERT INTO users (FirstName, LastName,Email, Mobile, FatherName, MotherName )VALUES ('Ashley', 'Nil', 'AshleySweta@gmail.com', '8940056941','Sweta','Ash')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println("Yay, values added!")
	results, err := db.Query("SELECT FirstName FROM users")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var outputDb OutputDb
		err = results.Scan(&outputDb.name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(outputDb.name)
	}
}
