package main

import (
	"flag"
	"fmt"

	"moneyrate/routes"

	"github.com/gin-gonic/gin"
)

var serverPort string

func main() {
	flag.StringVar(&serverPort, "p", "9000", "server port")
	flag.Parse()

	fmt.Println("", serverPort)

	router := gin.Default()
	routes.Attach(router)

	addr := ":" + serverPort
	router.Run(addr)
}
