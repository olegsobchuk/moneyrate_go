package controllers

import (
	"net/http"

	"moneyrate/config"
	"moneyrate/model"

	"github.com/gin-gonic/gin"
)

// Root page (welcome page)
func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome!",
	})
}

// Login provides data for Login page
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"user":    model.User{},
		"message": "Sucess!",
	})
}

// SessionCreate checks and authorize user
func SessionCreate(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Not found"})
}
