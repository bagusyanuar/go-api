package database

import "go-api/src/model"

type MemberToUser struct {
	model.Member
	User model.User `gorm:"foreignKey:UserID"`
}

type AdminToUser struct {
	model.Admin
	User model.User `gorm:"foreignKey:UserID"`
}
