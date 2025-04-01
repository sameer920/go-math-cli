package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition", "+"},
	Short:   "Add 2 numbers",
	Long:    "Carry out addition on two numbers.",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := Add(args[0], args[1])
		if err == nil {
			fmt.Printf("Addition of %s and %s = %s \n\n", args[0], args[1], result)
		} else if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
