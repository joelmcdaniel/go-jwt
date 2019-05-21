package main

import (
	"database/sql"
	"jwt-auth-restapi/controllers"
	"jwt-auth-restapi/driver"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/protected", controller.TokenVerifyMiddleWare(controller.ProtectedEndpoint(db))).Methods("GET")
	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")

	log.Println("Listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
