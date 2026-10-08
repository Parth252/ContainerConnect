package ecshlp

import (
	"context"
	"fmt"
)

func ExecuteCommandInContainer(ctx context.Context, cluster string, task string, container string, command []string) error {
	discovery, err := LoadDiscovery(ctx)
	if err != nil {
		return err
	}

	tasks, err := discovery.ListTasks(ctx, cluster)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		return fmt.Errorf("no ECS tasks found in cluster %s", cluster)
	}

	if task == "" {
		task = tasks[0] // Default to the first task if not specified
	}

	fmt.Printf("Executing command in container %s of task %s in cluster %s\n", container, task, cluster)

	// Here you would implement the logic to execute the command in the specified container.
	// This is a placeholder for demonstration purposes.
	fmt.Printf("Command to execute: %v\n", command)

	return nil
}
