package users

import (
	"fmt"
	"goAuthService/utils"
	"regexp"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByEmail(email string) *User {
	db := utils.GetDB()
	user := &User{}
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		fmt.Println("User find error")
		return nil
	}
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

func (credentials *Credentials) AreValid() bool {
	user := GetUserByEmail(credentials.Email)
	if user == nil {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	return !(err != nil && err == bcrypt.ErrMismatchedHashAndPassword)
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
