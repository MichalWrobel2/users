package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDBConnection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	name := "DATABASE"
	user := os.Getenv(name + "_USER")
	host := os.Getenv(name + "_HOST")
	port := os.Getenv(name + "_PORT")
	dbst := os.Getenv(name + "_NAME")
	pass := os.Getenv(name + "_PASSWORD")
	logmode, errLogMode := strconv.ParseBool(os.Getenv(name + "_LOGMODE"))
	if errLogMode != nil {
		fmt.Print("ENV DB Logmode error")
	}
	databaseConnectionInfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable host=%s port=%s password=%s", user, dbst, host, port, pass)
	db, err := gorm.Open("postgres", databaseConnectionInfo)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Cannot connect to database with config: ", databaseConnectionInfo)
	}
	db.LogMode(logmode)
	DB = db
}

func GetDB() *gorm.DB {
	if DB == nil {
		InitDBConnection()
		defer DB.Close()
	}
	return DB
}
