package controllers

import (
	"net/http"

	"moneyrate/config"
	"moneyrate/model"
	"moneyrate/service/token"

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

	config.DB.Where("email = ?", user.Email).First(&user)
	if user.ID != 0 {
		jwt, err := token.New(map[string]interface{}{"userID": user.ID})
		if err != nil {
			panic("can't generate jwt token")
		}
		c.JSON(http.StatusOK, gin.H{"token": jwt})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Not found"})
}

// CheckToken - should be rewritten
func CheckToken(c *gin.Context) (int, error) {
	receivedToken := c.Request.Header["Token"][0]
	jwt, err := token.Parse(receivedToken)
	if err != nil {
		return 0, err
	}

	return jwt["userID"].(int), nil
}
