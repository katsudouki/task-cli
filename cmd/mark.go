package cmd

import (
	"fmt"
	"strconv"
	"task-cli/data"

	"github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
	Use:   "mark [status] [id]",
	Short: "Mark a task with 'in-progress' or 'done'",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		status := args[0]
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("ID invalid.")
			return
		}
		if err := data.MarkTask(id, status); err != nil {
			fmt.Println("Error on mark task:", err)
			return
		}
		fmt.Printf("Task marked with '%s'.\n", status)
	},
}

func init() {
	rootCmd.AddCommand(markCmd)
}
