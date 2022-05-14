package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// answerCmd represents the answer command
var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "Answer question",
	Run: func(cmd *cobra.Command, args []string) {
		// Read ID from the argument
		id, fault := cmd.Flags().GetString("id")

		// Get quiz details
		fmt.Print("Fetching questions from API...\n\n")
		quiz, fault := GetQuizDetails(id)

		if fault != nil {
			os.Stderr.WriteString(fault.Error())
			os.Exit(1)
		}

		defer quiz.Close()

		var details [3]QuizStructure
		fault = json.NewDecoder(quiz).Decode(&details)

		if fault != nil {
			os.Stderr.WriteString("Failed to parse question's JSON")
			os.Exit(1)
		}

		// go through the questions and output to the user with options

		var answer string
		var answers [3]AnswerResponse

		fmt.Print("Nathan Galea - Quiz starting\n")
		fmt.Print("----------------------------\n\n")

		for i, det := range details {

			// Show question
			fmt.Print("Question: " + det.Question + "\n\n")

			// Show answers
			for j, answer := range det.AnswersList {
				fmt.Print(strconv.Itoa(j+1) + ") " + answer + "\n")
			}

			fmt.Println("\nEnter your answer (ID)")

			// Read user input
			fmt.Scanln(&answer)
			fmt.Print("\n")

			// Validate user input
			answerNum, err := strconv.Atoi(answer)

			if err != nil {
				os.Stderr.WriteString("Answer must be a number")
				os.Exit(1)
			}

			answersLen := len(det.AnswersList)

			if answerNum <= 0 || answerNum > answersLen {
				answersLenStr := strconv.Itoa(answersLen)
				os.Stderr.WriteString("Answer must be in the range 1-" + answersLenStr)
				os.Exit(1)
			}

			answers[i] = AnswerResponse{
				AnswerID: answerNum - 1,
			}
		}

		// Send results to API
		answersResultReader, fault := PostCompletedQuiz(id, answers)

		if fault != nil {
			os.Stderr.WriteString("Failed to send answers. " + fault.Error())
			os.Exit(1)
		}

		defer answersResultReader.Close()

		respBody, fault := ioutil.ReadAll(answersResultReader)

		if fault != nil {
			os.Stderr.WriteString("Error reading response. " + fault.Error())
			os.Exit(1)
		}

		fmt.Println("Results:", string(respBody))
	},
}

// Send our answers to the API for processing
func PostCompletedQuiz(questionID string, answers Answers) (io.ReadCloser, error) {
	answersJSON, fault := json.Marshal(answers)

	if fault != nil {
		os.Stderr.WriteString("Failed to JSON encode answers. " + fault.Error())
		os.Exit(1)
	}

	address := connectionAPI + "/quiz/" + questionID + "/answers"
	resp, fault := http.Post(address, "application/json", bytes.NewBuffer(answersJSON))

	if fault != nil {
		return nil, fault
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP status is not OK")
	}

	return resp.Body, nil
}

func init() {
	rootCmd.AddCommand(answerCmd)
	answerCmd.Flags().StringP("id", "i", "", "Quiz ID")
}
