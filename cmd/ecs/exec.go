package ecs

import (
	"context"

	"github.com/Parth252/ContainerConnect/internal/ecshlp"
	"github.com/spf13/cobra"
)

var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in a running container in ECS",
	RunE: func(cmd *cobra.Command, args []string) error {
		return executeCommandInTask(
			cmd.Context(),
			args,
		)
	},
}

func RegisterExecCmd() {
	EcsCmd.AddCommand(ExecCmd)

	ExecCmd.Flags().StringVarP(&inputCluster, "cluster", "c", "", "ECS cluster name or ARN (required)")
	ExecCmd.Flags().StringVarP(&inputContainer, "container", "n", "", "ECS container name (optional, required if multiple containers exist in the task)")
	ExecCmd.Flags().StringVarP(&inputTask, "task", "t", "", "ECS task ID or ARN (required)")

	ExecCmd.MarkFlagRequired("cluster")
	ExecCmd.MarkFlagRequired("task")

}

// Implementations

func executeCommandInTask(ctx context.Context, command []string) error {
	execHelper := ecshlp.LoadExecHelper(inputCluster, inputTask, inputContainer)
	return execHelper.ExecuteCommandInTask(ctx, inputCluster, inputTask, inputContainer, command)
}
