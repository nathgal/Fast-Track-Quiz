package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/quiz/:id", GetQuiz)
	router.POST("/quiz/:id/answers", PostAnswer)
	router.Run("localhost:8080")

	log.Println("Started API server on port :8080")
}
