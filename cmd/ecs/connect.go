package ecs

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a running container in ECS",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connecting to ECS container...")
	},
}

func init() {
	EcsCmd.AddCommand(ConnectCmd)
}
