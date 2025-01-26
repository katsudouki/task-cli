package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Task CLI - simple task manager",
	Long:  `a simple task manager  written in go.`,
	Run: func(cmd *cobra.Command, args []string) {
		listCmd.Run(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
