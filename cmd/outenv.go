package cmd

import "github.com/spf13/cobra"

var outenvCmd = &cobra.Command{
	Use:   "outenv",
	Short: "Output environment variables",
	Long:  "Output environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newService()
		serviceCtx.OutEnv()
	},
}