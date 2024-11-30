package cmd

import (
	"fmt"
	"os"

	"github.com/justinbather/life/life/internal/config"
	"github.com/justinbather/life/life/internal/http"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "life",
	Short: "Track meals, workouts, and more, right from the terminal",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	config, err := config.ReadLifeConfig()
	if err != nil {
		fmt.Printf("Error reading Life Config: %s", err)
		os.Exit(1)
	}

	jwt, err := http.Authenticate(config)
	if err != nil {
		fmt.Printf("Error Authenticating: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Authenticated! \n %s\n", jwt)
	rootCmd.PersistentFlags().String("jwt", jwt, "")
}
