package main

import (
	// "fmt"
	// "reflect"
	"goAuthService/models"
	"goAuthService/utils"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	// "goAuthService/models"
	"goAuthService/controllers"
)

func main() {
	utils.InitDBConnection()
	models.Migrate()
	setupRouterAndServe()
}

func setupRouterAndServe() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.IndexHandler)
	router.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
	setupServer(router)
}

func setupServer(router *mux.Router) {
	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
