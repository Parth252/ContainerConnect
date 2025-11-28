package ecs

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "List ECS services in a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing ECS services...")
	},
}

func init() {
	EcsCmd.AddCommand(ListServicesCmd)
}
