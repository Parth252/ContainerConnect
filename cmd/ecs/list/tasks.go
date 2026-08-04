package list

import (
	"fmt"

	awshlp "github.com/Parth252/ContainerConnect/internal/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
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
		return listTasks(tasksCluster)
	},
}

func listTasks(cluster string) error {
	client, ctx, err := awshlp.LoadECSClient()
	if err != nil {
		return err
	}

	output, err := client.ListTasks(ctx, &ecs.ListTasksInput{Cluster: &cluster})
	if err != nil {
		return fmt.Errorf("failed to list ECS tasks: %w", err)
	}

	if len(output.TaskArns) == 0 {
		fmt.Println("No ECS tasks found in this cluster.")
		return nil
	}

	fmt.Println("ECS Tasks:")
	for _, arn := range output.TaskArns {
		fmt.Println(arn)
	}

	return nil
}

func RegisterTasks(parent *cobra.Command) {
	tasksCmd.Flags().StringVarP(&tasksCluster, "cluster", "c", "", "ECS cluster name or ARN (required)")
	parent.AddCommand(tasksCmd)
}
