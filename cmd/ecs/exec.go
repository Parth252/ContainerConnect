package ecs

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in a running container in ECS",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing command in ECS container...")
	},
}

func RegisterExecCmd(parent *cobra.Command) {
	parent.AddCommand(ExecCmd)
}
