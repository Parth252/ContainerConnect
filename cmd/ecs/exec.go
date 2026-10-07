package ecs

import (
	"context"
	"fmt"

	ecshlp "github.com/Parth252/ContainerConnect/internal/ecs"
	"github.com/spf13/cobra"
)

var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in a running container in ECS",
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeCommandInContainer(
			cmd.Context(),
			inputCluster,
			inputTask,
			inputContainer,
			args,
		)
	},
}

func RegisterExecCmd(parent *cobra.Command) {
	ExecCmd.Flags().StringVarP(&inputCluster, "cluster", "c", "", "ECS cluster name or ARN (required)")
	ExecCmd.Flags().StringVarP(&inputContainer, "container", "n", "", "ECS container name (required)")

	ExecCmd.MarkFlagRequired("cluster")
	ExecCmd.MarkFlagRequired("container")

	parent.AddCommand(ExecCmd)
}

// Implementations

func executeCommandInContainer(ctx context.Context, cluster string, task string, container string, command []string) error {
	discovery, err := ecshlp.LoadDiscovery(ctx)
	if err != nil {
		return err
	}

	tasks, err := discovery.ListTasks(ctx, cluster)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		return fmt.Errorf("no ECS tasks found in cluster %s", cluster)
	}

	if task == "" {
		task = tasks[0] // Default to the first task if not specified
	}

	fmt.Printf("Executing command in container %s of task %s in cluster %s\n", container, task, cluster)

	// Here you would implement the logic to execute the command in the specified container.
	// This is a placeholder for demonstration purposes.
	fmt.Printf("Command to execute: %v\n", command)

	return nil
}
