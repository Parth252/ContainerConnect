package ecs

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ListClustersCmd represents 'cc ecs list-clusters'
var ListClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "List ECS clusters in your AWS account",
	Run: func(cmd *cobra.Command, args []string) {
		// Placeholder for AWS SDK logic
		fmt.Println("Listing ECS clusters...")
	},
}

func init() {
	EcsCmd.AddCommand(ListClustersCmd)
}
