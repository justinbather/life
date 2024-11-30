/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/justinbather/life/life/internal/service"
	"github.com/justinbather/life/life/model"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

		workout := workoutFromFlags(cmd.Flags())
		jwt, _ := cmd.PersistentFlags().GetString("jwt")

		_, err := service.CreateWorkout(workout, jwt)
		if err != nil {
			fmt.Printf("Error Creating workout: %s", err)
			return nil
		}

		fmt.Println("Created workout successfully...")

		return nil
	},
}

func workoutFromFlags(flags *pflag.FlagSet) model.Workout {
	wType, _ := flags.GetString("type")
	wDur, _ := flags.GetInt("duration")
	wCals, _ := flags.GetInt("cals")
	wLoad, _ := flags.GetInt("load")
	wDesc, _ := flags.GetString("desc")
	user, _ := flags.GetString("user")

	workout := model.Workout{User: user, Type: wType, Duration: wDur, CaloriesBurned: wCals, Workload: wLoad, Description: wDesc}

	return workout
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
