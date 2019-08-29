package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone" `
	Role      string     `json:"role"`
	Propertys []Property `gorm:"foreignkey:UserRefer"`
}

type Property struct {
	gorm.Model
	LocationName string `json:"locationname"`
	Address      string `json:"address"`
	City         string `json:"city"`
	State        string `json:"state" `
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Photo        string `json:"photo"`
	UserRefer    uint
}

var propertys []Property
var users []User

//creates the initial migration for the User DB
func InitialMigration() {
	db, err := gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &Property{})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)
	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	// users = append(users, newUser)

	db, err = gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	fmt.Println(newUser)

	db.Create(&newUser)

	fmt.Println(w, "New User Successfully Created")
	json.NewEncoder(w).Encode(newUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	email := vars["email"]

	user := User{}
	db.Where("email = ?", email).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
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
func AllProperties(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var propertys []Property
	db.Find(&propertys)
	fmt.Println("{}", propertys)
	json.NewEncoder(w).Encode(propertys)
}

func NewProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProperty Property
	json.NewDecoder(r.Body).Decode(&newProperty)

	db, err = gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	fmt.Println(newProperty)

	db.Create(&newProperty)

	fmt.Println(w, "New User Successfully Created")
	json.NewEncoder(w).Encode(newProperty)
}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	email := vars["email"]

	property := Property{}
	db.Where("email = ?", email).Find(&property)
	db.Delete(&property)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func UpdateProperty(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// json.NewDecoder(r.Body).Decode(&newProperty)
	vars := mux.Vars(r)
	locationname := vars["firstname"]
	address := vars["lastname"]
	city := vars["email"]
	state := vars["role"]
	latitude := vars["role"]
	longitude := vars["role"]
	photo := vars["photo"]

	var property Property
	db.Where("locationname = ?", locationname).Find(&property)

	property.LocationName = locationname
	property.Address = address
	property.City = city
	property.State = state
	property.Latitude = latitude
	property.Longitude = longitude
	property.Photo = photo

	db.Save(&property)
	fmt.Fprintf(w, "Successfully Updated Property")
}