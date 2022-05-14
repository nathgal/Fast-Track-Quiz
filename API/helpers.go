package main

import (
	"log"
	"math"
	"strconv"
)

func GetSpecificQuiz(id string) *[3]Quiz {
	quizID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Responded with HTTP status 400 | Error with requested ID.")
		return nil
	}

	response, exist := QuizDataset[quizID]

	if !exist {
		log.Println("Responded with HTTP status 404 | Requested quiz ID was not found.")
		return nil
	}

	return &response
}

func GetResultsDifference(totalCorrect int, totalQuestions int) int {
	if totalCorrect == 0 {
		return 0
	}

	if totalCorrect == totalQuestions || totalTimesQuizTaken == 0 {
		return 100
	}

	prevLowerResults := 0

	for i := 0; i <= totalCorrect; i++ {
		prevLowerResults += correctlyAnswered[i]
	}

	if prevLowerResults == 0 {
		return 0
	}

	return int(math.Ceil(float64(prevLowerResults) / float64(totalTimesQuizTaken) * 100.0))
}
