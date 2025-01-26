package cmd

import (
	"fmt"
	"task-cli/data"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description] [priority]",
	Short: "Add new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		priority := args[0]
		description := args[1]
		id, err := data.AddTask(priority,description)
		if err != nil {
			fmt.Println("Error on create new task:", err)
			return
		}
		fmt.Printf("Task added successfully (ID: %d)\n", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
