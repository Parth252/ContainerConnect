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

func (e ExecHelper) ExecuteCommandInTask(ctx context.Context) error {
	discovery, err := LoadDiscovery(ctx)
	if err != nil {
		return fmt.Errorf("failed to load discovery: %w", err)
	}

	containers, err := discovery.ListContainers(ctx, e.inputCluster, e.inputTask)
	if err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	if len(containers) > 1 && e.inputContainer == "" {
		return fmt.Errorf("multiple containers found in task %s please specify a container with --container", e.inputTask)
	}

	taskDescription, err := discovery.DescribeTasks(ctx, e.inputCluster, []string{e.inputTask})
	if err != nil {
		return err
	}

	if len(taskDescription.Tasks) == 0 {
		return fmt.Errorf("task %s not found in cluster %s", e.inputTask, e.inputCluster)
	}

	task := taskDescription.Tasks[0]

	if !task.EnableExecuteCommand {
		return fmt.Errorf(
			"ECS Exec is not enabled for task %s",
			e.inputTask,
		)
	}

	return nil
}
