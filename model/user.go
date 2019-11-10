package model

import (
	"database/sql"

	"moneyrate/config"

	"github.com/jinzhu/gorm"
)

// User describes Users attributes
type User struct {
	gorm.Model
	Email                string `form:"email" json:"user" binding:"required"`
	Password             string `form:"password" json:"password" gorm:"-"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" gorm:"-"`
	EncPassword          *sql.NullString
}

// Create is a function which insert User into DB
func (u *User) Create() {
	config.DB.Create(u)
}
