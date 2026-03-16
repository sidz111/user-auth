package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"uniqueIndex" json:"user_id"`
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
