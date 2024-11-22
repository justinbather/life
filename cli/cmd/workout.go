/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/justinbather/life/cli/internal/service"
	"github.com/justinbather/life/cli/model"
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

		// fairly ugly, refetching these flags after validation
		wType, _ := cmd.Flags().GetString("type")
		wDur, _ := cmd.Flags().GetInt("duration")
		wCals, _ := cmd.Flags().GetInt("cals")
		wLoad, _ := cmd.Flags().GetInt("load")
		wDesc, _ := cmd.Flags().GetString("desc")
		user, _ := cmd.Flags().GetString("user")

		workout := model.Workout{User: user, Type: wType, Duration: wDur, CaloriesBurned: wCals, Workload: wLoad, Description: wDesc}

		workout, err = service.CreateWorkout(workout)
		if err != nil {
			return err
		}

		fmt.Println("Created workout successfully...")
		fmt.Printf("User: %s\nType: %s\nCalories Burned: %d\nDuration: %d\nWorkload: %d\nDescription: %s\n", workout.User, workout.Type, workout.CaloriesBurned, workout.Duration, workout.Workload, workout.Description)
		return nil
	},
}

func validateFlags(flags *pflag.FlagSet) error {
	wDur, err := flags.GetInt("duration")
	if err != nil {
		return err
	}
	if wDur == 0 {
		return fmt.Errorf("Duration is required. Must be greater than 0")
	}

	wCals, err := flags.GetInt("cals")
	if err != nil {
		return err
	}
	if wCals == 0 {
		return fmt.Errorf("Calories is required. Must be greater than 0")
	}
	wType, err := flags.GetString("type")
	if err != nil {
		return err
	}
	if wType == "" {
		return fmt.Errorf("Type is required")
	}

	_, err = flags.GetInt("load")
	if err != nil {
		return err
	}

	_, err = flags.GetString("desc")
	if err != nil {
		return err
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
