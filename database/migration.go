package database

import "go-api/src/model"

func Migrate() {
	CONFIG.AutoMigrate(&model.User{})
}
