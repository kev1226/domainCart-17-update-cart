package main

import (
	"log"
	"update-cart/config"
	"update-cart/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectRedis()

	r := gin.Default()
	routes.SetupRoutes(r)

	if err := r.Run(":3034"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
