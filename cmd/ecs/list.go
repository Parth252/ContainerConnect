package ecs

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Used to list ECS resources like clusters and services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ECS command selected — use a subcommand like 'list' or 'exec'")
	},
}

func init() {
	EcsCmd.AddCommand(listCmd)
}
