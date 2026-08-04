package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Used to list ECS resources like clusters, services, and tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list command selected — use a subcommand like 'clusters', 'services', or 'tasks'")
	},
}

func Register(parent *cobra.Command) {
	parent.AddCommand(ListCmd)
	RegisterClusters(ListCmd)
	RegisterServices(ListCmd)
	RegisterTasks(ListCmd)
}
