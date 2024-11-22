/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// workoutCmd represents the workout command
var workoutCmd = &cobra.Command{
	Use:   "workout",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Creating new workout")
		err := validateFlags(cmd.LocalFlags())
		if err != nil {
			return err
		}

		wType, _ := cmd.Flags().GetString("type")
		wDur, _ := cmd.Flags().GetInt("duration")
		wCals, _ := cmd.Flags().GetInt("cals")
		wLoad, _ := cmd.Flags().GetInt("load")
		wDesc, _ := cmd.Flags().GetString("desc")

		fmt.Println("Workout Details")
		fmt.Printf("Type: %s\nCalories Burned: %d\nDuration: %d\nWorkload: %d\nDescription: %s\n", wType, wCals, wDur, wLoad, wDesc)
		return nil
	},
}

func validateFlags(flags *pflag.FlagSet) error {
	wDur, _ := flags.GetInt("duration")
	if wDur == 0 {
		return fmt.Errorf("Duration is required. Must be greater than 0")
	}
	wCals, _ := flags.GetInt("cals")
	if wCals == 0 {
		return fmt.Errorf("Calories is required. Must be greater than 0")
	}
	wType, _ := flags.GetString("type")
	if wType == "" {
		return fmt.Errorf("Type is required")
	}

	return nil
}

func init() {
	newCmd.AddCommand(workoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//workoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	workoutCmd.Flags().String("type", "", "Declares the type of workout. Ex. Run, Weights. Required")
	workoutCmd.Flags().Int("duration", 0, "Declares the duration of the workout in minutes. Must be > 0")
	workoutCmd.Flags().Int("cals", 0, "Declares calories burned in workout. Must be > 0")
	workoutCmd.Flags().Int("load", 0, "Declares the workload of your workout, range from 0..10 (inclusive)")
	workoutCmd.Flags().String("desc", "", "An optional description for the workout")

	workoutCmd.MarkFlagRequired("type")
	workoutCmd.MarkFlagRequired("duration")
	workoutCmd.MarkFlagRequired("cals")
}
