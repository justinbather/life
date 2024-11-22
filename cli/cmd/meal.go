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

// mealCmd represents the meal command
var createMealCmd = &cobra.Command{
	Use:   "meal",
	Short: "Creates a meal with given information",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Creating new meal...")
		err := cmd.ValidateRequiredFlags()
		if err != nil {
			return err
		}

		meal := getMeal(cmd.Flags())

		created, err := service.CreateMeal(meal)
		if err != nil {
			fmt.Printf("Error creating meal: %s", err)
			return nil
		}

		fmt.Println("Created meal successfully")
		fmt.Printf("Type: %s\nCals: %d\nProtein: %d\nCarbs: %d\nFat: %d\nDescription: %s\n", created.Type, created.Calories, created.Protein, created.Carbs, created.Fat, created.Description)

		return nil
	},
}

func getMeal(flags *pflag.FlagSet) model.Meal {
	_type, _ := flags.GetString("type")
	cals, _ := flags.GetInt("cals")
	protein, _ := flags.GetInt("protein")
	carbs, _ := flags.GetInt("carbs")
	fat, _ := flags.GetInt("fat")
	desc, _ := flags.GetString("desc")

	meal := model.Meal{Type: _type, Calories: cals, Protein: protein, Carbs: carbs, Fat: fat, Description: desc}

	return meal
}

func init() {
	newCmd.AddCommand(createMealCmd)

	createMealCmd.Flags().String("type", "", "Required: Meal type ex. Lunch")
	createMealCmd.Flags().Int("cals", 0, "Required: Number of calories. Must be > 0")
	createMealCmd.Flags().Int("protein", 0, "Required: Protein in grams. Must be > 0")
	createMealCmd.Flags().Int("carbs", 0, "Required: Carbs in grams. Must be > 0")
	createMealCmd.Flags().Int("fat", 0, "Required: Fat in grams. Must be > 0")
	createMealCmd.Flags().String("desc", "", "Optional: A short description")

	createMealCmd.MarkFlagRequired("type")
	createMealCmd.MarkFlagRequired("cals")
	createMealCmd.MarkFlagRequired("protein")
	createMealCmd.MarkFlagRequired("carbs")
	createMealCmd.MarkFlagRequired("fat")
}
