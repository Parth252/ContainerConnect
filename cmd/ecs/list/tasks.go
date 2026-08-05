package list

import (
	"context"
	"fmt"

	ecshlp "github.com/Parth252/ContainerConnect/internal/ecs"
	"github.com/spf13/cobra"
)

var tasksCluster string

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List ECS tasks in a cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		if tasksCluster == "" {
			return fmt.Errorf("cluster name or ARN is required (--cluster)")
		}
		return listTasks(cmd.Context(), tasksCluster)
	},
}

func listTasks(ctx context.Context, cluster string) error {
	discovery, err := ecshlp.LoadDiscovery(ctx)
	if err != nil {
		return err
	}

	tasks, err := discovery.ListTasks(ctx, cluster)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No ECS tasks found in this cluster.")
		return nil
	}

	fmt.Println("ECS Tasks:")
	for _, arn := range tasks {
		fmt.Println(arn)
	}

	return nil
}

func RegisterTasks(parent *cobra.Command) {
	tasksCmd.Flags().StringVarP(&tasksCluster, "cluster", "c", "", "ECS cluster name or ARN (required)")
	parent.AddCommand(tasksCmd)
}
