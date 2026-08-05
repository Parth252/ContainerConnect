package ecs

import (
	"context"
	"fmt"

	awshlp "github.com/Parth252/ContainerConnect/internal/aws"
	awsecs "github.com/aws/aws-sdk-go-v2/service/ecs"
)

// API is the subset of the ECS client needed to discover ECS resources.
type API interface {
	ListClusters(context.Context, *awsecs.ListClustersInput, ...func(*awsecs.Options)) (*awsecs.ListClustersOutput, error)
	ListServices(context.Context, *awsecs.ListServicesInput, ...func(*awsecs.Options)) (*awsecs.ListServicesOutput, error)
	ListTasks(context.Context, *awsecs.ListTasksInput, ...func(*awsecs.Options)) (*awsecs.ListTasksOutput, error)
}

// Discovery provides ECS resource discovery operations.
type Discovery struct {
	client API
}

func LoadDiscovery(ctx context.Context) (*Discovery, error) {
	client, err := awshlp.LoadECSClient(ctx)
	if err != nil {
		return nil, err
	}

	return NewDiscovery(client), nil
}

func NewDiscovery(client API) *Discovery {
	return &Discovery{client: client}
}

func (d *Discovery) ListClusters(ctx context.Context) ([]string, error) {
	output, err := d.client.ListClusters(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("list ECS clusters: %w", err)
	}

	return output.ClusterArns, nil
}

func (d *Discovery) ListServices(ctx context.Context, cluster string) ([]string, error) {
	output, err := d.client.ListServices(ctx, &awsecs.ListServicesInput{Cluster: &cluster})
	if err != nil {
		return nil, fmt.Errorf("list ECS services: %w", err)
	}

	return output.ServiceArns, nil
}

func (d *Discovery) ListTasks(ctx context.Context, cluster string) ([]string, error) {
	output, err := d.client.ListTasks(ctx, &awsecs.ListTasksInput{Cluster: &cluster})
	if err != nil {
		return nil, fmt.Errorf("list ECS tasks: %w", err)
	}

	return output.TaskArns, nil
}
