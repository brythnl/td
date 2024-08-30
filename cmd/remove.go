package cmd

import (
	"log"
	"slices"

	"github.com/brythnl/td/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a task",
	Long:    `Remove a task`,
	Run:     runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

// removeTasks removes tasks of the passed in positions
func removeTasks(tasks []todo.Task, positions []int) []todo.Task {
	return slices.DeleteFunc(tasks, func(t todo.Task) bool {
		for _, p := range positions {
			if t.Position == p {
				return true
			}
		}
		return false
	})
}

func runRemove(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	tasks, err := todo.ReadTasks(dataFile)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	if len(args) < 1 {
		log.Fatalln("Provide at least one task (number) to remove")
	}

	tasks = removeTasks(tasks, argsToPositions(args, len(tasks)))
	todo.OrderPositions(tasks)

	err = todo.WriteTasks(dataFile, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	showTasks(tasks, true)
}
