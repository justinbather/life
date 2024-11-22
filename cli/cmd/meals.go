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

// mealsCmd represents the meals command
var mealsCmd = &cobra.Command{
	Use:   "meals",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		flagVal, _ := cmd.Flags().GetString("timeframe")
		tf, err := timeframe.ParseTimeframe(flagVal)
		if err != nil {
			return err
		}

		fmt.Printf("Fetching meals from the past %s...\n", tf.String())

		user, _ := cmd.Flags().GetString("user")

		mp := tf.GetRange()

		meals, err := service.GetMeals(user, mp)
		if err != nil {
			return err
		}

		fmt.Printf("Fetched %d meals successfully...\n", len(meals))
		fmt.Println("Results")
		fmt.Println("#   Type    Cals     Protein    Carbs    Fat    Desc")
		for idx, m := range meals {
			fmt.Printf("%d   %s    %d         %d      %d      %d     %s\n", idx, m.Type, m.Calories, m.Protein, m.Carbs, m.Fat, m.Description)
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(mealsCmd)

	mealsCmd.Flags().String("timeframe", "week", "Optional: Query timeframe. Defaults to week")
}
