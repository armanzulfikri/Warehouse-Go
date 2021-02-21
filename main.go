package main

import (
	"fmt"
	"os"
	"time"
	"warehouse/config"
	"warehouse/database/migrations"
	"warehouse/database/seeder"
	"warehouse/web"

	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 500 * time.Hour,
	}))
	router.Run(":" + os.Getenv("PORT"))
}
