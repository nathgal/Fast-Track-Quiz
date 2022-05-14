package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// [Description] : Get Quiz functionality by ID
func GetQuiz(c *gin.Context) {
	log.Println("New Request | Class: main.go, Function: GetQuestions - ID:")

	response := GetSpecificQuiz(c.Param("id"))

	if response == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Requested question was not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
	return
}

// [Description] : Save answers
func PostAnswer(c *gin.Context) {
	log.Println("New Request | Class: main.go, Function: Post - ID:")

	quiz := GetSpecificQuiz(c.Param("id"))

	if err := c.BindJSON(&results); err != nil {
		log.Println("Responded with HTTP status 400 | JSON was not binded correctly")
		return
	}

	// Check if length of results are equal to questions
	totalQuestions := len(quiz)

	if totalQuestions != len(results) {
		log.Println("Responded with HTTP status 400 | Incorrect number of results compared to questions")
		return
	}

	// Calculate total correct answers
	correctAnswers := 0
	for questionIx, answer := range results {
		if answer.AnswerID == quiz[questionIx].CorrectAnswer {
			correctAnswers++
		}
	}

	resultPercentage := GetResultsDifference(correctAnswers, totalQuestions)

	// Update general statistics
	correctlyAnswered[correctAnswers]++
	totalTimesQuizTaken++

	var res = struct {
		CorrectAnswers    int    `json:"correctAnswers"`
		ComparingToOthers string `json:"comparingToOthers"`
	}{
		CorrectAnswers:    correctAnswers,
		ComparingToOthers: "You were better than " + strconv.Itoa(resultPercentage) + "% of all quizzers",
	}

	log.Println("Responded with HTTP status 200 OK")
	c.IndentedJSON(http.StatusOK, res)
	return
}
