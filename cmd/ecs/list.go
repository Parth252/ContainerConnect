package ecs

import (
	"context"
	"fmt"

	ecshlp "github.com/Parth252/ContainerConnect/internal/ecs"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Used to list ECS resources like clusters, services, and tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list command selected — use a subcommand like 'clusters', 'services', or 'tasks'")
	},
}

func Register(parent *cobra.Command) {
	parent.AddCommand(ListCmd)
	RegisterClusters(ListCmd)
	RegisterServices(ListCmd)
	RegisterTasks(ListCmd)
}

// Implementations

//Clusters: List all ECS clusters in the configured AWS account

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

//Services: List all ECS services in a cluster

var servicesCluster string

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "List ECS services in a cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		if servicesCluster == "" {
			return fmt.Errorf("cluster name or ARN is required (--cluster)")
		}
		return listServices(cmd.Context(), servicesCluster)
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
	servicesCmd.Flags().StringVarP(&servicesCluster, "cluster", "c", "", "ECS cluster name or ARN (required)")
	parent.AddCommand(servicesCmd)
}

//Tasks: List all ECS tasks in a cluster

var tasksCluster string

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List ECS tasks in a cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		if tasksCluster == "" {
			return fmt.Errorf("cluster name or ARN is required (--cluster)")
		}
		return listTasks(cmd.Context(), tasksCluster)
	},
}

func listTasks(ctx context.Context, cluster string) error {
	discovery, err := ecshlp.LoadDiscovery(ctx)
	if err != nil {
		return err
	}

	tasks, err := discovery.ListTasks(ctx, cluster)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No ECS tasks found in this cluster.")
		return nil
	}

	fmt.Println("ECS Tasks:")
	for _, arn := range tasks {
		fmt.Println(arn)
	}

	return nil
}

func RegisterTasks(parent *cobra.Command) {
	tasksCmd.Flags().StringVarP(&tasksCluster, "cluster", "c", "", "ECS cluster name or ARN (required)")
	parent.AddCommand(tasksCmd)
}
