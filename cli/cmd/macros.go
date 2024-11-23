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

// macrosCmd represents the macros command
var macrosCmd = &cobra.Command{
	Use:   "macros",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		tfVal, _ := cmd.Flags().GetString("timeframe")
		tf, err := timeframe.ParseTimeframe(tfVal)
		if err != nil {
			return err
		}

		fmt.Printf("Fetching macros from the past %s...\n", tf.String())

		user, _ := cmd.Flags().GetString("user")

		mp := tf.GetRange()

		macros, err := service.GetMacros(user, mp)
		if err != nil {
			return err
		}

		fmt.Println("Fetched macros successfully...")
		fmt.Println(macros)

		return nil
	},
}

func init() {
	getCmd.AddCommand(macrosCmd)

	macrosCmd.Flags().String("timeframe", "week", "Optional: Specifies timeframe of the query. Default to week. Options: today|week|month|year")
	macrosCmd.Flags().String("view", "basic", "Optional: Specifies view type. Defaults to basic. Options: basic|full")
}
