package main

import (
	"fmt"
	"warehouse/config"
	"warehouse/database/migrations"
	"warehouse/database/seeder"
	"warehouse/web"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found ..")
	}
}

func main() {
	config.Connect()
	//panggil seeder dan migrate disini
	migrations.Migrations()
	seeder.Seeder()
	router := web.Route(gin.Default())
	router.Run(":8080")
}
