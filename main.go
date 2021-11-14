package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"quiz-app/config/database"
	"quiz-app/routes"
)

func main() {
	_ = godotenv.Load()

	r := gin.Default()

	routes.Route(r)

	if err := database.Connect(); err != nil {
		fmt.Println(err)
		return
	}

	_ = r.Run(":" + os.Getenv("PORT"))
}