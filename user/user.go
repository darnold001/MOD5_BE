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
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Photo        string `json:"photo"`
	UserRefer    uint   `json:"userID"`
	Accountings []Accounting `gorm:"foreignkey:PropertyRefer"`
}

type Accounting struct{
	gorm.Model
	Rent string `json:"rent"`
	RentTotal float64 `json:"renttotal"`
	Expenses string `json:"expenses"`
	ExpenseTotal float64 `json:"expensetotal"`
	PL float64 `json:"pl"`
	Year string `json:"Year"`
	PropertyRefer uint `json:"propertyID"`
}

var propertys []Property
var users []User
var accountings []Accounting

//creates the initial migration for the User DB
func InitialMigration() {
	db, err := gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &Property{}, &Accounting{})
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

func SelectUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "user=Arnold dbname=user sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	email := vars["email"]

	user := User{}
	db.Where("email = ?", email).Find(&user)
	json.NewEncoder(w).Encode(&user)
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
	id := vars["id"]

	property := Property{}
	db.Where("id = ?", id).Find(&property)
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
	locationname := vars["locationname"]
	address := vars["address"]
	latitude := vars["latitude"]
	longitude := vars["longitude"]
	photo := vars["photo"]

	var property Property
	db.Where("locationname = ?", locationname).Find(&property)

	property.LocationName = locationname
	property.Address = address
	property.Latitude = latitude
	property.Longitude = longitude
	property.Photo = photo

	db.Save(&property)
	fmt.Fprintf(w, "Successfully Updated Property")
}
