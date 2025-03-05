package main

import (
	"log"

	"aas.dev/internal/config"
	"aas.dev/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	r := gin.Default()

	routes.SetupUserRoutes(r)
	routes.SetupAdminRoutes(r)
	routes.SetupMainRoutes(r)
	routes.SetupGraphRoutes(r)

	log.Println("server running on port: ", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
