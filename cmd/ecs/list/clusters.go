package list

import (
	"context"
	"fmt"

	awshlp "github.com/Parth252/ContainerConnect/internal/aws"
	"github.com/spf13/cobra"
)

var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "List all ECS clusters in the configured AWS account",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listClusters()
	},
}

func listClusters() error {
	fmt.Println("Fetching ECS clusters...")

	ctx := context.Background()
	client, err := awshlp.LoadECSClient(ctx)
	if err != nil {
		return err
	}

	output, err := client.ListClusters(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to list ECS clusters: %w", err)
	}

	if len(output.ClusterArns) == 0 {
		fmt.Println("No ECS clusters found in this AWS account/region.")
		return nil
	}

	fmt.Println("\nECS Clusters:")
	fmt.Println("-------------------------")
	for _, arn := range output.ClusterArns {
		fmt.Printf("- %s\n", arn)
	}
	fmt.Println("-------------------------")

	return nil
}

func RegisterClusters(parent *cobra.Command) {
	parent.AddCommand(clustersCmd)
}
