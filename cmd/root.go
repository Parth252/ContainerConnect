package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cc",
	Short: "ContainerConnect - connect to running containers in ECS/EKS",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running root command")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
