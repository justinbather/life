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

// mealsCmd represents the meals command
var mealsCmd = &cobra.Command{
	Use:   "meals",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		tfVal, _ := cmd.Flags().GetString("timeframe")
		tf, err := timeframe.ParseTimeframe(tfVal)
		if err != nil {
			return err
		}

		user, _ := cmd.Flags().GetString("user")

		mp := tf.GetRange()

		meals, err := service.GetMeals(user, mp)
		if err != nil {
			return err
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		t.AppendHeader(table.Row{"Date", "Type", "Calories", "Protein (g)", "Carbs (g)", "Fat (g)", "Description"})

		for _, m := range meals {
			t.AppendRow(table.Row{m.Date.Format(time.DateOnly), m.Type, m.Calories, m.Protein, m.Carbs, m.Fat, m.Description})
			t.AppendSeparator()
		}

		t.SortBy([]table.SortBy{{Name: "Date", Mode: table.Asc}})

		t.Render()

		return nil
	},
}

func init() {
	getCmd.AddCommand(mealsCmd)

	mealsCmd.Flags().String("timeframe", "week", "Optional: Query timeframe. Defaults to week")
}
