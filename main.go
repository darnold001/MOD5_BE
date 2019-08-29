package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darnold001/rentalapi/user"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go ORM")
	user.InitialMigration()

	handleRequests()
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", user.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", user.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", user.UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/users", user.NewUser).Methods("POST")
	myRouter.HandleFunc("/properties", user.AllProperties).Methods("GET")
	myRouter.HandleFunc("/property/{name}", user.DeleteProperty).Methods("DELETE")
	myRouter.HandleFunc("/property/{name}/{email}", user.UpdateProperty).Methods("PUT")
	myRouter.HandleFunc("/properties", user.NewProperty).Methods("POST")
	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(myRouter)))

}
