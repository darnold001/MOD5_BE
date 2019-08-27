package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	FirstName string
	LastName string
	Email string
	Role string
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "user.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}


func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "user.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")

	db, err := gorm.Open("sqlite3", "user.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	firstname := vars["firstname"]
	lastname := vars["lastname"]
	email := vars["email"]
	role := vars["role"]

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(role)
	fmt.Println(email)

	db.Create(&User{FirstName: firstname, LastName: lastname, Email: email, Role: role})
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "user.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	email := vars["email"]

	var user User
	db.Where("email = ?", email).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "user.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	firstname := vars["firstname"]
	lastname := vars["lastname"]
	email := vars["email"]
	role := vars["role"]

	var user User
	db.Where("email = ?", email).Find(&user)

	user.FirstName = firstname
	user.LastName = lastname
	user.Email = email
	user.Role = role

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}