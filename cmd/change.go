package cmd

import (
	"github.com/spf13/cobra"
)

// changeCmd represents the change command
var changeCmd = &cobra.Command{
	Use:     "change",
	Aliases: []string{"c"},
	Short:   "Change current project",
	Long:    `Change current project`,
	Run:     runChange,
}

func init() {
	rootCmd.AddCommand(changeCmd)
}

func runChange(cmd *cobra.Command, args []string) {
	// TODO: implement
	// - changing projects (also persist change in config file)
	// - add new project tasks file when one from argument doesn't exist
}
