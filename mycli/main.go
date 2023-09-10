package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "greet",
		Short: "A CLI to demonstrate Cobra usage",
		Long:  `This CLI demonstrate how to create a CLI application with cobra.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Greetings,%s!\n", args[0])
		},
	}

	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "Prints a 'Hello, world!' message",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, world!")
		},
	}

	rootCmd.AddCommand(helloCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
