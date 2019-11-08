package main

import (
	"flag"
	"time"

	"moneyrate/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var serverPort string

func main() {
	flag.StringVar(&serverPort, "p", "9000", "server port")
	flag.Parse()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))
	routes.Attach(router)

	addr := ":" + serverPort
	router.Run(addr)
}
