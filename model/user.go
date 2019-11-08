package model

import "time"

// User describes Users attributes
type User struct {
	Email                string `form:"email" json:"user" binding:"required"`
	Password             string `form:"password" json:"password"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation"`
	encPassword          string
	UpdatedAt            time.Time `form:"updated_at" json:"updated_at"`
	CreatedAt            time.Time `form:"created_at" json:"created_at"`
}
