package list

import (
	"context"
	"fmt"

	ecshlp "github.com/Parth252/ContainerConnect/internal/ecs"
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
		return listServices(cmd.Context(), cluster)
	},
}

func listServices(ctx context.Context, cluster string) error {
	discovery, err := ecshlp.LoadDiscovery(ctx)
	if err != nil {
		return err
	}

	services, err := discovery.ListServices(ctx, cluster)
	if err != nil {
		return err
	}

	if len(services) == 0 {
		fmt.Println("No ECS services found in this cluster.")
		return nil
	}

	fmt.Println("ECS Services:")
	for _, arn := range services {
		fmt.Println(arn)
	}

	return nil
}

func RegisterServices(parent *cobra.Command) {
	servicesCmd.Flags().StringVarP(&cluster, "cluster", "c", "", "ECS cluster name or ARN (required)")
	parent.AddCommand(servicesCmd)
}
