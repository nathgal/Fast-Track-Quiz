package main

type Quiz struct {
	Question      string    `json:"question"`
	AnswersList   [3]string `json:"answersList"`
	CorrectAnswer int       `json:"correctAnswer"`
}

var results []struct {
	AnswerID int `json:"answerId"`
}

// In-memory data - QuizDataset

var QuizDataset = map[int][3]Quiz{
	1: [3]Quiz{
		Quiz{
			Question:      "Who was Darth Vader?",
			AnswersList:   [3]string{"Anakin Skywalker", "Yoda", "Mace Windu"},
			CorrectAnswer: 0,
		},
		Quiz{
			Question:      "Which character showed in person in Better Call Saul and not in Breaking Bad?",
			AnswersList:   [3]string{"Saul Goodman", "Nacho Vargo", "Gus Fring"},
			CorrectAnswer: 1,
		},
		Quiz{
			Question:      "How many Spider-man films were there post 2000?",
			AnswersList:   [3]string{"3", "4", "5"},
			CorrectAnswer: 2,
		},
	},
	2: [3]Quiz{
		Quiz{
			Question:      "What is the back-end system for this test program written in?",
			AnswersList:   [3]string{"Golang", ".NET", "PHP"},
			CorrectAnswer: 0,
		},
		Quiz{
			Question:      "What role within Fast Track is this test requested for?",
			AnswersList:   [3]string{"Front-end Developer", "Back-end Developer", "BI Analyst"},
			CorrectAnswer: 1,
		},
		Quiz{
			Question:      "How much productivity is increaed up to when using Fast Track's CRM?",
			AnswersList:   [3]string{"40%", "50%", "60%"},
			CorrectAnswer: 2,
		},
	},
}

var correctlyAnswered [4]int // The number of times users correctly answered the questions
var totalTimesQuizTaken = 0  // Total times the quiz was taken
