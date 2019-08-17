package controllers

import (
	"encoding/json"
	"goAuthService/models"
	"goAuthService/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	if !IsValidUser(user, w) {
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)
	user.Create()
}

func IsValidUser(user models.User, w http.ResponseWriter) bool {
	if len(user.FirstName) < 1 || len(user.LastName) < 1 || len(user.Email) < 1 || len(user.Password) < 1 {
		utils.RespondJSON(w, http.StatusBadRequest, "Bad request")
		return false
	}

	if !user.IsFirstNameValid() {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid first name")
		return false
	}

	if !user.IsLastNameValid() {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid last name")
		return false
	}

	if !user.IsEmailValid() {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid email")
		return false
	}

	if !user.IsPasswordValid() {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid password")
		return false
	}

	return true
}
