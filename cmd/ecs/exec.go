package ecs

import (
	"context"

	ecshlp "github.com/Parth252/ContainerConnect/internal/ecshlp"
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
	return ecshlp.ExecuteCommandInContainer(ctx, cluster, task, container, command)
}
