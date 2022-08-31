package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

var DB_USER string
var DB_PASS string
var Users []User

func envVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Not able to find .env file")
	}
	return os.Getenv(key)
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All user Endpoint hit")
	db, err := gorm.Open("mysql", "root:Abcdefghij2124$@tcp(127.0.0.1:3306)/users?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	// db := dbConnect(DB_USER, DB_PASS)
	db.Find(&Users)
	fmt.Printf("Users detail : %+v", Users)
	json.NewEncoder(w).Encode(Users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New user Endpoint Hit")
	db, err := gorm.Open("mysql", "root:Abcdefghij2124$@tcp(127.0.0.1:3306)/users")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	Users = append(Users, user)
	json.NewEncoder(w).Encode(user)
	db.Create(&user)
	fmt.Printf("New user : %+v", Users)
	fmt.Fprintf(w, "New user successfully created")

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user Endpoint Hit")
	db, err := gorm.Open("mysql", "root:Abcdefghij2124$@tcp(127.0.0.1:3306)/users?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	var user User
	db.Where("name=?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "Successfully user deleted")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update user Endpoint Hit")
	db, err := gorm.Open("mysql", "root:Abcdefghij2124$@tcp(127.0.0.1:3306)/users?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	vars := mux.Vars(r)
	id := vars["id"]
	var user User
	db.Where("id = ?", id).Find(&user)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updateUserField User
	json.Unmarshal(reqBody, &updateUserField)
	fmt.Printf("Update user : %+v", updateUserField)
	user.Name = updateUserField.Name
	user.Email = updateUserField.Email
	db.Save(&user)
	fmt.Fprintf(w, "Update completed")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}
func initialMigration(DB_USER string, DB_PASS string) {
	db, err := gorm.Open("mysql", DB_USER+":"+DB_PASS+"@tcp(127.0.0.1:3306)/users")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	fmt.Println("Connected to DB")

	// Migrate the schema
	db.AutoMigrate(&User{})
}

// func dbConnect(DB_USER string, DB_PASS string) *mysqlConn {
// 	db, err := gorm.Open("mysql", DB_USER+":"+DB_PASS+"@tcp(127.0.0.1:3306)/users?parseTime=true")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer db.Close()
// 	return db
// }

func main() {
	fmt.Println("GO ORM tutorial")
	DB_USER = envVariable("DB_USER")
	DB_PASS = envVariable("DB_PASS")
	initialMigration(DB_USER, DB_PASS)
	// Handle Subsequent requests
	handleRequests()
}
