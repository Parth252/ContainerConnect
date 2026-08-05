package list

import (
	"context"
	"fmt"

	ecshlp "github.com/Parth252/ContainerConnect/internal/ecs"
	"github.com/spf13/cobra"
)

var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "List all ECS clusters in the configured AWS account",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listClusters(cmd.Context())
	},
}

func listClusters(ctx context.Context) error {
	fmt.Println("Fetching ECS clusters...")

	discovery, err := ecshlp.LoadDiscovery(ctx)
	if err != nil {
		return err
	}

	clusters, err := discovery.ListClusters(ctx)
	if err != nil {
		return err
	}

	if len(clusters) == 0 {
		fmt.Println("No ECS clusters found in this AWS account/region.")
		return nil
	}

	fmt.Println("\nECS Clusters:")
	fmt.Println("-------------------------")
	for _, arn := range clusters {
		fmt.Printf("- %s\n", arn)
	}
	fmt.Println("-------------------------")

	return nil
}

func RegisterClusters(parent *cobra.Command) {
	parent.AddCommand(clustersCmd)
}
