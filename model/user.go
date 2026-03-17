package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string `gorm:"type:varchar(36);uniqueIndex" json:"uuid"`
	Username string `gorm:"type:varchar(255);uniqueIndex" json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
