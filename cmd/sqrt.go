package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sqrtCmd = &cobra.Command{
	Use:     "sqrt",
	Aliases: []string{"Sqrt", "rt"},
	Short:   "Square root of given number",
	Long:    "Calculates the square root of the given number",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := Sqrt(args[0])
		if err == nil {
			fmt.Printf("Square root of %s = %s \n\n", args[0], result)
		} else if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sqrtCmd)
}
