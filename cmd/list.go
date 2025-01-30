package cmd

import (
	"fmt"
	"task-cli/data"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	Green   = color.New(color.FgGreen).SprintFunc()
	Red     = color.New(color.FgRed).SprintFunc()
	Yellow  = color.New(color.FgYellow).SprintFunc()
	Blue    = color.New(color.FgBlue).SprintFunc()
	Magenta = color.New(color.FgMagenta).SprintFunc()
	Reset   = color.New(color.Reset).SprintFunc()
)

var listCmd = &cobra.Command{
	Use:   "list [status]",
	Short: "List tasks",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		status := ""
		if len(args) > 0 {
			status = args[0]
		}

		tasks, err := data.ListTasks(status)
		if err != nil {
			fmt.Println("Error on list tasks:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("Empty tasks list.")
			return
		}

		for _, task := range tasks {
			taskStatus := "[ ]"
			if task.Status == "todo" {
				taskStatus = "[ ]"
			} else if task.Status == "in-progress" {
				taskStatus = "[...]"
			} else {
				taskStatus = "[✔]"
			}

			var colorpr func(a ...interface{}) string
			if task.Priority == "low" {
				colorpr = Green
			} else if task.Priority == "normal" {
				colorpr = Blue
			} else {
				colorpr = Red
			}

			fmt.Printf("(%d) %s %s [%s] %s\n",
				task.ID,
				Green(taskStatus),
				Reset(""),
				colorpr(task.Priority),
				Magenta(task.Description),
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
