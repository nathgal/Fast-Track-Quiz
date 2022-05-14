package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var connectionAPI = "http://localhost:8080"

var rootCmd = &cobra.Command{
	Use:   "CLI",
	Short: "Short description: Quick quiz applciation for Fast Track",
	Long:  "Long description: Quick quiz with a few questions and a few alternatives for each question using Golang as the backend.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
