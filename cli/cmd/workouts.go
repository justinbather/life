/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
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
		tfVal, _ := cmd.Flags().GetString("timeframe")
		tf, err := timeframe.ParseTimeframe(tfVal)
		if err != nil {
			return err
		}

		user, _ := cmd.Flags().GetString("user")

		dateMap := tf.GetRange()

		workouts, err := service.GetWorkouts(user, dateMap)
		if err != nil {
			return err
		}

		if len(workouts) == 0 {
			fmt.Println("No Workouts found.")
			return nil
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Date", "Type", "Calories", "Intensity", "Duration (mins)", "Description"})

		for _, w := range workouts {
			t.AppendRow(table.Row{w.Date.Format(time.DateOnly), w.Type, w.CaloriesBurned, w.Workload, w.Duration, w.Description})
			t.AppendSeparator()
		}

		t.SortBy([]table.SortBy{{Name: "Date", Mode: table.Asc}})

		t.Render()

		return nil
	},
}

func init() {
	getCmd.AddCommand(workoutsCmd)
	// $ life get workouts --timeframe=week|month|year|today
	workoutsCmd.Flags().String("timeframe", "week", "Optional: Specify timeframe of query. Defaults to week. Options: today, week, month, year")
}
