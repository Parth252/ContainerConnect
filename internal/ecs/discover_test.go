package ecs

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsecs "github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

type fakeClient struct {
	clusters *awsecs.ListClustersOutput
	services *awsecs.ListServicesOutput
	tasks    *awsecs.ListTasksOutput
	describe *awsecs.DescribeTasksOutput
	err      error
}

func (f *fakeClient) ListClusters(context.Context, *awsecs.ListClustersInput, ...func(*awsecs.Options)) (*awsecs.ListClustersOutput, error) {
	return f.clusters, f.err
}

func (f *fakeClient) ListServices(context.Context, *awsecs.ListServicesInput, ...func(*awsecs.Options)) (*awsecs.ListServicesOutput, error) {
	return f.services, f.err
}

func (f *fakeClient) ListTasks(context.Context, *awsecs.ListTasksInput, ...func(*awsecs.Options)) (*awsecs.ListTasksOutput, error) {
	return f.tasks, f.err
}

func (f *fakeClient) DescribeTasks(context.Context, *awsecs.DescribeTasksInput, ...func(*awsecs.Options)) (*awsecs.DescribeTasksOutput, error) {
	return f.describe, f.err
}

func TestDiscoveryListsResources(t *testing.T) {
	discovery := NewDiscovery(&fakeClient{
		clusters: &awsecs.ListClustersOutput{ClusterArns: []string{"cluster-1"}},
		services: &awsecs.ListServicesOutput{ServiceArns: []string{"service-1"}},
		tasks:    &awsecs.ListTasksOutput{TaskArns: []string{"task-1"}},
	})

	clusters, err := discovery.ListClusters(context.Background())
	if err != nil || !reflect.DeepEqual(clusters, []string{"cluster-1"}) {
		t.Fatalf("ListClusters() = %v, %v", clusters, err)
	}

	services, err := discovery.ListServices(context.Background(), "cluster-1")
	if err != nil || !reflect.DeepEqual(services, []string{"service-1"}) {
		t.Fatalf("ListServices() = %v, %v", services, err)
	}

	tasks, err := discovery.ListTasks(context.Background(), "cluster-1")
	if err != nil || !reflect.DeepEqual(tasks, []string{"task-1"}) {
		t.Fatalf("ListTasks() = %v, %v", tasks, err)
	}
}

func TestDiscoveryListsContainersInTask(t *testing.T) {
	discovery := NewDiscovery(&fakeClient{
		describe: &awsecs.DescribeTasksOutput{
			Tasks: []types.Task{{
				Containers: []types.Container{
					{Name: aws.String("application")},
					{Name: aws.String("sidecar")},
				},
			}},
		},
	})

	containers, err := discovery.ListContainers(context.Background(), "cluster-1", "task-1")
	if err != nil || !reflect.DeepEqual(containers, []string{"application", "sidecar"}) {
		t.Fatalf("ListContainers() = %v, %v", containers, err)
	}
}

func TestDiscoveryReturnsDescribeTaskFailures(t *testing.T) {
	discovery := NewDiscovery(&fakeClient{
		describe: &awsecs.DescribeTasksOutput{
			Failures: []types.Failure{{Reason: aws.String("MISSING")}},
		},
	})

	_, err := discovery.ListContainers(context.Background(), "cluster-1", "missing-task")
	if err == nil || err.Error() != "describe ECS task \"missing-task\": MISSING" {
		t.Fatalf("expected task failure error, got %v", err)
	}
}

func TestDiscoveryWrapsClientErrors(t *testing.T) {
	clientErr := errors.New("AWS unavailable")
	discovery := NewDiscovery(&fakeClient{err: clientErr})

	_, err := discovery.ListClusters(context.Background())
	if err == nil || !errors.Is(err, clientErr) {
		t.Fatalf("expected wrapped client error, got %v", err)
	}
}
