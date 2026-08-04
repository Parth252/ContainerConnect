package ecs

import "testing"

func TestEcsCommandRegistersListCommand(t *testing.T) {
	command, _, err := EcsCmd.Find([]string{"list"})
	if err != nil {
		t.Fatalf("expected list command to be registered: %v", err)
	}
	if command == nil || command.Name() != "list" {
		t.Fatal("expected list command to be registered under ecs")
	}
}
