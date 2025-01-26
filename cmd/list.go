package cmd

import (
	"fmt"
	"task-cli/data"

	"github.com/spf13/cobra"
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
			taskstatus := "[ ]"
			if task.Status == "todo" {
				taskstatus = "[ ]"
			} else if task.Status == "in-progress"{
                taskstatus="[...]"
			} else {
                taskstatus ="[✔]"
			}
            colorpr :=""
			if task.Priority == "low"{
				colorpr=data.Green
			} else if task.Priority =="normal"{
                colorpr=data.Blue
			} else{
                colorpr=data.Red
			}




			fmt.Printf("(%d)"+ data.Green +" %s"+ data.Reset +"  ["+colorpr+"%s"+data.Reset +"]"+data.Magenta+" %s"+data.Reset+"\n", task.ID, taskstatus,task.Priority, task.Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
