package cmd

import (
	"fmt"
	"strconv"
	"task-cli/data"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID invalid.")
			return
		}
		if err := data.DeleteTask(id); err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}
		fmt.Println("Task deleted.")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
