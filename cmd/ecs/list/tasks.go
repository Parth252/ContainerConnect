package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List ECS tasks in a cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("list tasks is not implemented yet")
		return nil
	},
}

func RegisterTasks(parent *cobra.Command) {
	parent.AddCommand(tasksCmd)
}
