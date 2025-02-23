package main

import (
	"log"

	"aas.dev/internal/config"
	"aas.dev/pkg/routes"
	"aas.dev/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	db := utils.ConnectDB(cfg)
	r := gin.Default()

	routes.SetupUserRoutes(r, db)
	routes.SetupMainRoutes(r)

	log.Println("server running on port: ", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
