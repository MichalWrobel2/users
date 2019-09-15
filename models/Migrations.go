package models

import (
	"goAuthService/models/users"
	"goAuthService/utils"
)

func Migrate() {
	db := utils.GetDB()
	db.AutoMigrate(
		&users.User{},
	)
}
