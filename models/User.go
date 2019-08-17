package models

import (
	"fmt"
	"goAuthService/utils"
	"regexp"
)

type User struct {
	ID        uint64 `json:"id" gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique"`
	Password  string
}

func (user *User) Get() *User {
	db := utils.GetDB()
	fmt.Println(db.Find(user))
	return user
}

func (user *User) Create() {
	db := utils.GetDB()
	fmt.Println(db.Find(user))
	fmt.Println(db.Create(user))
}

func (user *User) IsFirstNameValid() bool {
	fname, _ := regexp.MatchString(".{3,}", user.FirstName)
	return fname
}

func (user *User) IsLastNameValid() bool {
	lname, _ := regexp.MatchString(".{3,}", user.LastName)
	return lname
}

func (user *User) IsEmailValid() bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(user.Email)
}
func (user *User) IsPasswordValid() bool {
	specialChar, _ := regexp.MatchString("[!@#$%^&*(){}\"|:;'\\\\[\\].,~`?/+_=-]+", user.Password)
	downcaseLetter, _ := regexp.MatchString("[a-z]+", user.Password)
	uppercaseLetter, _ := regexp.MatchString("[A-Z]+", user.Password)
	number, _ := regexp.MatchString("\\d", user.Password)
	length, _ := regexp.MatchString(".{8,}", user.Password)

	return (specialChar &&
		downcaseLetter &&
		uppercaseLetter &&
		number &&
		length)
}
