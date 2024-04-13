package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/shoetan/ecom/controllers"
	"github.com/shoetan/ecom/db"
)


func main() {
	
	//start database connection
	db,err := db.Db()

	if err != nil {
		log.Fatal("Can't connect to the database:", err)
	}
	//handlers that handles the request
	router:=mux.NewRouter()

	router.HandleFunc("/createUser", controllers.CreateUser(db) ).Methods("POST")
	router.HandleFunc("/getUsers", controllers.GetAllUsers(db) ).Methods("GET")
	router.HandleFunc("/loginUser", controllers.LoginUser(db) ).Methods("POST")

	fmt.Println("Server is running on :5355")



  log.Fatal(http.ListenAndServe(":5355", router))
}
