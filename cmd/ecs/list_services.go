package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/spf13/cobra"
)

var cluster string

var listServicesCmd = &cobra.Command{
	Use:   "services",
	Short: "List ECS services in a cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cluster == "" {
			return fmt.Errorf("cluster name or ARN is required (--cluster)")
		}
		return listServices(cluster)
	},
}

func listServices(cluster string) error {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	client := ecs.NewFromConfig(cfg)

	output, err := client.ListServices(ctx, &ecs.ListServicesInput{
		Cluster: &cluster,
	})
	if err != nil {
		return fmt.Errorf("failed to list ECS services: %w", err)
	}

	if len(output.ServiceArns) == 0 {
		fmt.Println("No ECS services found in this cluster.")
		return nil
	}

	fmt.Println("ECS Services:")
	for _, arn := range output.ServiceArns {
		fmt.Println(arn)
	}

	return nil
}

func init() {
	listServicesCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "ECS cluster name or ARN (required)")
	listCmd.AddCommand(listServicesCmd)
}
