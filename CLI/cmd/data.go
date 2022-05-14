package cmd

type QuizStructure struct {
	Question    string    `json:"question"`
	AnswersList [3]string `json:"answersList"`
}

type AnswerResponse struct {
	AnswerID int `json:"answerId"`
}

type Answers [3]AnswerResponse
