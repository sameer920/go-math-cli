package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var squareCmd = &cobra.Command{
	Use:     "square",
	Aliases: []string{"sq"},
	Short:   "Square a number.",
	Long:    "Returns the square of the given number.",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := Square(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Square of %s = %s \n\n", args[0], result)
	},
}

func init() {
	rootCmd.AddCommand(squareCmd)
}
