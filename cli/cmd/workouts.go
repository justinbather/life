/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/justinbather/life/cli/internal/service"
	"github.com/justinbather/life/cli/pkg/timeframe"
	"github.com/spf13/cobra"
)

// workoutsCmd represents the workouts command
var workoutsCmd = &cobra.Command{
	Use:   "workouts",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		flagVal, _ := cmd.Flags().GetString("timeframe")

		tf, err := timeframe.ParseTimeframe(flagVal)
		if err != nil {
			return err
		}

		user, _ := cmd.Flags().GetString("user")
		dateMap := tf.GetRange()

		fmt.Println("Fetching workouts...")

		workouts, err := service.GetWorkouts(user, dateMap)
		if err != nil {
			return err
		}

		if len(workouts) == 0 {
			fmt.Println("No Workouts found.")
			return nil
		}

		fmt.Println("Fetched workouts successfully...")
		fmt.Println("Results:")
		fmt.Println("#      Type       Cals      Description")
		for idx, w := range workouts {
			fmt.Printf("%d      %s       %d      %s\n", idx, w.Type, w.CaloriesBurned, w.Description)
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(workoutsCmd)
	// $ life get workouts --timeframe=week|month|year|today
	workoutsCmd.Flags().String("timeframe", "week", "Optional: Specify timeframe of query. Defaults to week. Options: today, week, month, year")
}
