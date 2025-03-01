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
	Run:   runPrune,
}

func init() {
	rootCmd.AddCommand(pruneCmd)
}

func runPrune(cmd *cobra.Command, args []string) {
	wpName, wpFile, err := td.GetWorkingProject()
	if err != nil {
		log.Fatalf("unable to get current working project: %v\n", err)
	}
	tasks, err := td.ReadTasks(wpFile)
	if err != nil {
		log.Fatalf("unable to read tasks: %v\n", err)
	}

	var checkedPositions []int
	for _, t := range tasks {
		if t.Checked {
			checkedPositions = append(checkedPositions, t.Position)
		}
	}

	tasks = removeTasks(tasks, checkedPositions)
	td.OrderPositions(tasks)

	err = td.WriteTasks(wpFile, tasks)
	if err != nil {
		log.Fatalf("unable to write tasks: %v\n", err)
	}

	td.PrintHeader(wpName)
	td.PrintTasks(tasks, td.ShowUnchecked)
}
