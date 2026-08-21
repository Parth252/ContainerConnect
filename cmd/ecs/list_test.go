package ecs

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestRegisterAddsExpectedSubcommands(t *testing.T) {
	parent := &cobra.Command{Use: "ecs"}
	RegisterListCmd(parent)

	var listCmd *cobra.Command
	for _, cmd := range parent.Commands() {
		if cmd.Name() == "list" {
			listCmd = cmd
			break
		}
	}

	if listCmd == nil {
		t.Fatal("expected list subcommand to be registered")
	}

	registered := map[string]bool{}
	for _, cmd := range listCmd.Commands() {
		registered[cmd.Name()] = true
	}

	for _, name := range []string{"clusters", "services", "tasks"} {
		if !registered[name] {
			t.Fatalf("expected %q subcommand to be registered under list", name)
		}
	}
}
