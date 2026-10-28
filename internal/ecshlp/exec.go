package ecshlp

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

type ExecHelper struct {
	inputCluster   string
	inputTask      string
	inputContainer string
}

func LoadExecHelper(inputCluster, inputTask, inputContainer string) *ExecHelper {
	return &ExecHelper{
		inputCluster:   inputCluster,
		inputTask:      inputTask,
		inputContainer: inputContainer,
	}
}

func (e ExecHelper) ExecuteCommandInTask(ctx context.Context) error {
	discovery, err := LoadDiscovery(ctx)
	if err != nil {
		return fmt.Errorf("failed to load discovery: %w", err)
	}

	// Validate that the task exists and that the container is specified if there are multiple containers in the task
	containers, err := discovery.ListContainers(ctx, e.inputCluster, e.inputTask)
	if err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	if len(containers) > 1 && e.inputContainer == "" {
		return fmt.Errorf("multiple containers found in task %s please specify a container with --container, the available containers are: %v", e.inputTask, containers)
	} else if e.inputContainer == "" {
		e.inputContainer = containers[0]
	}

	taskDescriptions, err := discovery.DescribeTasks(ctx, e.inputCluster, []string{e.inputTask})
	if err != nil {
		return err
	}

	// Check if ECS Exec is enabled for the task and if the ECS Exec agent is running for the container
	taskDescription := taskDescriptions.Tasks[0]

	if !taskDescription.EnableExecuteCommand {
		return fmt.Errorf(
			"ECS Exec is not enabled for task %s",
			e.inputTask,
		)
	}

	for _, container := range taskDescription.Containers {
		if aws.ToString(container.Name) != e.inputContainer {
			continue
		}

		for _, agent := range container.ManagedAgents {
			if agent.Name == ecsTypes.ManagedAgentNameExecuteCommandAgent {
				if *agent.LastStatus != "RUNNING" {
					return fmt.Errorf(
						"ECS Exec agent is not running for container %s (status: %s)",
						e.inputContainer,
						*agent.LastStatus,
					)
				}

				return nil
			}
		}

		return fmt.Errorf("ECS Exec agent not found for container %s", e.inputContainer)
	}

	return fmt.Errorf("container %s not found in task %s", e.inputContainer, e.inputTask)
}
