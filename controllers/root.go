package controllers

import (
	"net/http"

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
