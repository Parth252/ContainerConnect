package cmd

import (
	"fmt"

	"github.com/Parth252/ContainerConnect/cmd/ecs"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cc",
	Short: "ContainerConnect - connect to running containers in ECS/EKS",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running root command")
	},
}

func Execute() {
	RootCmd.AddCommand(ecs.EcsCmd)
	cobra.CheckErr(RootCmd.Execute())
}
