package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	awsecs "github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/spf13/cobra"
)

var listClustersCmd = &cobra.Command{
	Use:   "list-clusters",
	Short: "List all ECS clusters in the configured AWS account",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listClusters()
	},
}

func listClusters() error {
	fmt.Println("Fetching ECS clusters...")

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	client := awsecs.NewFromConfig(cfg)

	output, err := client.ListClusters(ctx, &awsecs.ListClustersInput{})
	if err != nil {
		return fmt.Errorf("failed to list ECS clusters: %w", err)
	}

	if len(output.ClusterArns) == 0 {
		fmt.Println("No ECS clusters found in this AWS account/region.")
		return nil
	}

	fmt.Println("\n ECS Clusters:")
	fmt.Println("-------------------------")
	for _, arn := range output.ClusterArns {
		fmt.Printf("- %s\n", arn)
	}
	fmt.Println("-------------------------")

	return nil
}

func init() {
	EcsCmd.AddCommand(listClustersCmd)
}
