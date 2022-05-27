package database

import "go-api/src/model"

func Migrate() {
	CONFIG.AutoMigrate(&model.User{})
	CONFIG.AutoMigrate(&model.Member{})
	CONFIG.AutoMigrate(&model.Admin{})
	CONFIG.AutoMigrate(&model.Subject{})

	//create relation
	CONFIG.AutoMigrate(&MemberToUser{})
	CONFIG.AutoMigrate(&AdminToUser{})
}
