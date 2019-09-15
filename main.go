package main

import (
	"goAuthService/controllers"
	"goAuthService/models"
	"goAuthService/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
	setupServer(router)
}

func setupServer(router *mux.Router) {
	address := os.Getenv("HOST")
	server := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
