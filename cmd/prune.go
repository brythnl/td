package cmd

import (
	"log"

	"github.com/brythnl/td/td"
	"github.com/spf13/cobra"
)

// pruneCmd represents the prune command
var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove all checked tasks",
	Long:  `Remove all checked tasks`,
	Run:   runPrune,
}

func init() {
	rootCmd.AddCommand(pruneCmd)
}

func runPrune(cmd *cobra.Command, args []string) {
	projectName, projectFile := td.GetProject()
	tasks, err := td.ReadTasks(projectFile)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	var checkedPositions []int
	for _, t := range tasks {
		if t.Checked {
			checkedPositions = append(checkedPositions, t.Position)
		}
	}

	tasks = removeTasks(tasks, checkedPositions)
	td.OrderPositions(tasks)

	err = td.WriteTasks(projectFile, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	td.ShowTasks(tasks, td.ShowUnchecked, projectName)
}
