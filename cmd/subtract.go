package cmd

import (
	"fmt"

	"strconv"

	"github.com/spf13/cobra"
)

var subtractCmd = &cobra.Command{
	Use:     "subtract",
	Aliases: []string{"subtraction", "sub"},
	Short:   "Subtract one number from another.",
	Long:    "Given two integers, this command will subtract the second integer from the first.",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		temp, err := Subtract(args[0], args[1])
		if err != nil {
			fmt.Println(err)
			return
		} else {
			difference, err := strconv.ParseFloat(temp, 64)
			if err == nil {
				result := int(difference)
				fmt.Printf("Subtracting %s from %s = %v \n\n", args[1], args[0], result)
			} else {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)
}
