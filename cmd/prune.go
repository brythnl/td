package cmd

import (
	"log"

	"github.com/brythnl/td/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	dataFile := viper.GetString("datafile")
	tasks, err := todo.ReadTasks(dataFile)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	var checkedPositions []int
	for _, t := range tasks {
		if t.Checked {
			checkedPositions = append(checkedPositions, t.Position)
		}
	}

	tasks = RemoveTasks(tasks, checkedPositions)
	todo.OrderPositions(tasks)

	err = todo.WriteTasks(dataFile, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	showTasks(tasks, false)
}
