/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/justinbather/life/cli/internal/service"
	"github.com/spf13/cobra"
)

var ERR_INVALID_TIMEFRAME error = errors.New("Invalid timeframe.")

type timeframe int

func (t timeframe) String() string {
	return timeframeStrings[t]
}

func (t timeframe) GetRange() map[string]string {

	now := time.Now()
	eod := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Format(time.RFC3339)

	switch t {
	case TODAY:
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Format(time.RFC3339)

		return map[string]string{"start": start, "end": eod}
	case WEEK:
		sevenDaysAgo := now.AddDate(0, 0, -7).Format(time.RFC3339)
		return map[string]string{"start": sevenDaysAgo, "end": eod}
	case MONTH:
		monthAgo := now.AddDate(0, -1, 0).Format(time.RFC3339)
		return map[string]string{"start": monthAgo, "end": eod}
	case YEAR:
		yearAgo := now.AddDate(-1, 0, 0).Format(time.RFC3339)
		return map[string]string{"start": yearAgo, "end": eod}
	}

	return nil
}

const (
	UNKNOWN timeframe = iota
	TODAY
	WEEK
	MONTH
	YEAR
)

var timeframeStrings = map[timeframe]string{
	UNKNOWN: "unknown",
	TODAY:   "today",
	WEEK:    "week",
	MONTH:   "month",
	YEAR:    "year",
}

// workoutsCmd represents the workouts command
var workoutsCmd = &cobra.Command{
	Use:   "workouts",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		flagVal, _ := cmd.Flags().GetString("timeframe")

		tf, err := parseTimeframe(flagVal)
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

func parseTimeframe(val string) (timeframe, error) {
	switch strings.ToLower(val) {
	case "today":
		return TODAY, nil
	case "week":
		return WEEK, nil
	case "month":
		return MONTH, nil
	case "year":
		return YEAR, nil
	default:
		return UNKNOWN, ERR_INVALID_TIMEFRAME
	}
}

func init() {
	getCmd.AddCommand(workoutsCmd)
	// $ life get workouts --timeframe=week|month|year|today
	workoutsCmd.Flags().String("timeframe", "week", "Optional: Specify timeframe of query. Defaults to week. Options: today, week, month, year")
}
