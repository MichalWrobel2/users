package models

import (
	"fmt"
	"goAuthService/utils"
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        uint64    `json:"id" gorm:"AUTO_INCREMENT"`
	UUID      uuid.UUID `json:"uuid" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" gorm:"unique"`
	Password  string
	//Roles     []string
}

func (user *User) Get() *User {
	db := utils.GetDB()
	fmt.Println(db.Find(user))
	return user
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("UUID", uuid)
}
func (user *User) Create() {
	db := utils.GetDB()
	db.Create(user)
}

func (user *User) IsFirstNameValid() bool {
	IsFirstNameValid, _ := regexp.MatchString(".{3,}", user.FirstName)
	return IsFirstNameValid
}

func (user *User) IsLastNameValid() bool {
	IsLastNameValid, _ := regexp.MatchString(".{3,}", user.LastName)
	return IsLastNameValid
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
