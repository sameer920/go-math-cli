package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "math",
	Short: "go-math is a cli tool for performing basic mathematical operations",
	Long:  "go-math is a cli tool for performing basic mathematical operations - addition, multiplication, division, square, square root and subtraction.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			// Print an error message and show help
			fmt.Println("Error: No arguments provided. You must specify an operation (e.g., add, subtract, multiply, etc.)")
			cmd.Help() // Show the help text
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Math '%s'\n", err)
		os.Exit(1)
	}
}
