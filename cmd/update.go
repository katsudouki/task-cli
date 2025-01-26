package cmd

import (
	"fmt"
	"strconv"
	"task-cli/data"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [new priority] [new description]",
	Short: "update a task",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID invalid.")
			return
		}
		newPriority := args[1] 
		newDescription := args[2]
		if err := data.UpdateTask(id,newPriority, newDescription); err != nil {
			fmt.Println("Error on update task:", err)
			return
		}
		fmt.Println("Task updated")

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
