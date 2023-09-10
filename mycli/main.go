package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "greet",
		Short: "A CLI to demonstrate Cobra usage",
		Long:  `This CLI demonstrate how to create a CLI application with cobra.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Greetings,%s!\n", args[1])
		},
	}
}
