package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	FirstName string  `json:"firstname"`
	LastName  string  `json:"lastname"`
	Email     string  `json:"email"`
	Role      string `json:"role"`
}

var users []User

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
	w.Header().Set("Content-Type", "application/json")
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	// users = append(users, newUser)

	db, err = gorm.Open("sqlite3", "user.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	
	fmt.Println(newUser)


	db.Create(&newUser)


	fmt.Println(w, "New User Successfully Created")
	json.NewEncoder(w).Encode(newUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "user.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	email := vars["email"]

	user:=User{}
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
