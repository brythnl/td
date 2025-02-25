package cmd

import (
	"log"

	"github.com/brythnl/td/todo"
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
	project := todo.GetProjectFile()
	tasks, err := todo.ReadTasks(project)
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
	todo.OrderPositions(tasks)

	err = todo.WriteTasks(project, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	todo.ShowTasks(tasks, todo.ShowUnchecked)
}
