package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"warehouse/web"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found ..")
	}
}

func main()  {
	router := web.Route(gin.Default())
	router.Run(":8080")
}
