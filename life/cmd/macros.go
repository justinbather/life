/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/justinbather/life/life/internal/service"
	"github.com/justinbather/life/life/pkg/timeframe"
	"github.com/spf13/cobra"
)

// macrosCmd represents the macros command
var macrosCmd = &cobra.Command{
	Use:   "macros",
	Short: "retrieves your aggregated history of meals and workouts",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		tfVal, _ := cmd.Flags().GetString("timeframe")
		tf, err := timeframe.ParseTimeframe(tfVal)
		if err != nil {
			return err
		}

		user, _ := cmd.Flags().GetString("user")
		jwt, _ := cmd.PersistentFlags().GetString("jwt")

		mp := tf.GetRange()

		macros, err := service.GetMacros(user, mp, jwt)
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Date", "Calories", "Protein (g)", "Carbs (g)", "Fat (g)", "Calories Burned", "# of workouts"})

		totalCalsIn := 0
		totalCalsOut := 0
		totalWorkouts := 0

		for _, m := range macros {
			totalCalsOut += m.CalsBurned
			totalCalsIn += m.CalsIn
			totalWorkouts += m.Workouts

			t.AppendRow(table.Row{m.Date.Format(time.DateOnly), m.CalsIn, m.Protein, m.Carbs, m.Fat, m.CalsBurned, m.Workouts})
			t.AppendSeparator()
		}

		t.AppendRow(table.Row{"Totals", "Calories (net)", "# of workouts"})
		t.AppendFooter(table.Row{"", totalCalsIn - totalCalsOut, totalWorkouts})

		t.SortBy([]table.SortBy{{Name: "Date", Mode: table.Asc}})

		t.Render()

		return nil
	},
}

func init() {
	getCmd.AddCommand(macrosCmd)

	macrosCmd.Flags().String("timeframe", "week", "Optional: Specifies timeframe of the query. Default to week. Options: today|week|month|year")
	macrosCmd.Flags().String("view", "basic", "Optional: Specifies view type. Defaults to basic. Options: basic|full")
}
