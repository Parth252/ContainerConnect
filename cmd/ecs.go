package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ecsCmd represents the "ecs" command under root
var ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "Commands related to AWS ECS clusters and services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ECS command selected — use a subcommand like 'list' or 'exec'")
	},
}

func init() {
	rootCmd.AddCommand(ecsCmd)
}
