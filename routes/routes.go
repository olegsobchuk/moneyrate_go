package routes

import (
	"moneyrate/controllers"

	"github.com/gin-gonic/gin"
)

// Attach run existing routes
func Attach(router *gin.Engine) {
	router.GET("/", controllers.Root)
	router.GET("/login", controllers.Login)
}
