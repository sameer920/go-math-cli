package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var percentCmd = &cobra.Command{
	Use:     "percent",
	Aliases: []string{"%", "perc"},
	Short:   "Returns the given percentage of the given number.",
	Long:    "Given a number and the percentage to calculate from that number, this command will calcualte that percentage and return it.",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := Percentage(args[0], args[1])
		if err == nil {
			fmt.Printf("%s%% of %s = %s \n\n", args[1], args[0], result)
		} else if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(percentCmd)
}
