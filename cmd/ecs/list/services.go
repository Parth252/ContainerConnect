package list

import (
	"context"
	"fmt"

	awshlp "github.com/Parth252/ContainerConnect/internal/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/spf13/cobra"
)

var cluster string

var servicesCmd = &cobra.Command{
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
	client, err := awshlp.LoadECSClient(ctx)
	if err != nil {
		return err
	}

	output, err := client.ListServices(ctx, &ecs.ListServicesInput{Cluster: &cluster})
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

func RegisterServices(parent *cobra.Command) {
	servicesCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "ECS cluster name or ARN (required)")
	parent.AddCommand(servicesCmd)
}
