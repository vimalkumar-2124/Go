package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type OutputDb struct {
	name string `json:"name"`
}

func envVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Not able to find .env file")
	}
	return os.Getenv(key)
}
func main() {
	dbUser := envVariable("DB_USER")
	dbPass := envVariable("DB_PASS")
	// fmt.Printf("Pass : %s", dbUser)
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/students")
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
