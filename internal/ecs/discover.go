package ecs

import (
	"context"
	"fmt"

	awshlp "github.com/Parth252/ContainerConnect/internal/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsecs "github.com/aws/aws-sdk-go-v2/service/ecs"
)

// API is the subset of the ECS client needed to discover ECS resources.
type API interface {
	ListClusters(context.Context, *awsecs.ListClustersInput, ...func(*awsecs.Options)) (*awsecs.ListClustersOutput, error)
	ListServices(context.Context, *awsecs.ListServicesInput, ...func(*awsecs.Options)) (*awsecs.ListServicesOutput, error)
	ListTasks(context.Context, *awsecs.ListTasksInput, ...func(*awsecs.Options)) (*awsecs.ListTasksOutput, error)
	DescribeTasks(context.Context, *awsecs.DescribeTasksInput, ...func(*awsecs.Options)) (*awsecs.DescribeTasksOutput, error)
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

// ListContainers returns the names of every container in an ECS task.
func (d *Discovery) ListContainers(ctx context.Context, cluster, task string) ([]string, error) {
	output, err := d.client.DescribeTasks(ctx, &awsecs.DescribeTasksInput{
		Cluster: &cluster,
		Tasks:   []string{task},
	})
	if err != nil {
		return nil, fmt.Errorf("describe ECS task: %w", err)
	}

	if len(output.Failures) > 0 {
		failure := output.Failures[0]
		return nil, fmt.Errorf("describe ECS task %q: %s", task, aws.ToString(failure.Reason))
	}
	if len(output.Tasks) == 0 {
		return nil, fmt.Errorf("describe ECS task %q: task was not found", task)
	}

	containers := make([]string, 0, len(output.Tasks[0].Containers))
	for _, container := range output.Tasks[0].Containers {
		containers = append(containers, aws.ToString(container.Name))
	}

	return containers, nil
}
