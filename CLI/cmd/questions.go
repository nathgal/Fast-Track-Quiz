package cmd

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func GetQuizDetails(id string) (io.ReadCloser, error) {

	// Validate requested ID
	validID, fault := strconv.Atoi(id)

	if fault != nil || validID <= 0 {
		return nil, errors.New("Selected quiz number is not valid")
	}

	// Get quiz details
	address := connectionAPI + "/quiz/" + id
	resp, fault := http.Get(address)

	if fault != nil {
		return nil, fault
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP status is not OK")
	}

	return resp.Body, nil
}

var startQuizCommand = &cobra.Command{
	Use:   "startQuiz",
	Short: "Retrieve quiz from API server.",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		details, fault := GetQuizDetails(id)

		if fault != nil {
			os.Stderr.WriteString("Failed to retrieve the quiz.")
			os.Exit(1)
		}

		defer details.Close()
		respBody, fault := ioutil.ReadAll(details)

		if fault != nil {
			os.Stderr.WriteString("Error reading response. " + fault.Error())
			os.Exit(1)
		}

		fmt.Println("Response:\n", string(respBody))
	},
}

func init() {
	rootCmd.AddCommand(startQuizCommand)
	startQuizCommand.Flags().StringP("id", "i", "", "Quiz ID")
}
