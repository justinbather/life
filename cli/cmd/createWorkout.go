/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/justinbather/life/cli/internal/service"
	"github.com/justinbather/life/cli/model"
	"github.com/spf13/cobra"
)

// workoutCmd represents the workout command
var createWorkoutCmd = &cobra.Command{
	Use:   "workout",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Creating new workout...")
		if err := cmd.ValidateRequiredFlags(); err != nil {
			return err
		}

		wType, _ := cmd.Flags().GetString("type")
		wDur, _ := cmd.Flags().GetInt("duration")
		wCals, _ := cmd.Flags().GetInt("cals")
		wLoad, _ := cmd.Flags().GetInt("load")
		wDesc, _ := cmd.Flags().GetString("desc")
		user, _ := cmd.Flags().GetString("user")

		workout := model.Workout{User: user, Type: wType, Duration: wDur, CaloriesBurned: wCals, Workload: wLoad, Description: wDesc}

		w, err := service.CreateWorkout(workout)
		if err != nil {
			fmt.Printf("Error Creating workout: %s", err)
			return nil
		}

		fmt.Println("Created workout successfully...")
		fmt.Printf("User: %s\nType: %s\nCalories Burned: %d\nDuration: %d\nWorkload: %d\nDescription: %s\n", w.User, w.Type, w.CaloriesBurned, w.Duration, w.Workload, w.Description)
		return nil
	},
}

func init() {
	newCmd.AddCommand(createWorkoutCmd)
	createWorkoutCmd.Flags().String("type", "", "Declares the type of workout. Ex. Run, Weights. Required")
	createWorkoutCmd.Flags().Int("duration", 0, "Declares the duration of the workout in minutes. Must be > 0")
	createWorkoutCmd.Flags().Int("cals", 0, "Declares calories burned in workout. Must be > 0")
	createWorkoutCmd.Flags().Int("load", 0, "Declares the workload of your workout, range from 0..10 (inclusive)")
	createWorkoutCmd.Flags().String("desc", "", "An optional description for the workout")

	createWorkoutCmd.MarkFlagRequired("type")
	createWorkoutCmd.MarkFlagRequired("duration")
	createWorkoutCmd.MarkFlagRequired("cals")
}
