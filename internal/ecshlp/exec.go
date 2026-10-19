package ecshlp

import (
	"context"
	"fmt"
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

func (e ExecHelper) ExecuteCommandInTask(ctx context.Context, cluster string, task string, container string, command []string) error {
	discovery, err := LoadDiscovery(ctx)
	if err != nil {
		return fmt.Errorf("failed to load discovery: %w", err)
	}

	containers, err := discovery.ListContainers(ctx, cluster, task)
	if err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	if len(containers) > 1 && container == "" {
		return fmt.Errorf("multiple containers found in task %s please specify a container with --container", task)
	}

	fmt.Printf("Command to execute: %v\n", command)

	return nil
}
