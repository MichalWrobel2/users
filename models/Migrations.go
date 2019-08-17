package models

import (
	"goAuthService/utils"
)

func Migrate() {
	db := utils.GetDB()
	db.AutoMigrate(
		&User{},
	)
}
