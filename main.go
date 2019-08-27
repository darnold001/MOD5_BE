package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World")
}


func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}



func main(){
	fmt.Println("Go ORM")

	initialMigration()
	
	HandleRequests()
}