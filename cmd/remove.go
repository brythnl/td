package cmd

import (
	"fmt"
	"log"
	"slices"

	"github.com/brythnl/td/td"
	"github.com/spf13/cobra"
)

var removeAllOpt bool

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a task",
	Run:     runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolVarP(&removeAllOpt, "all", "a", false, "all tasks")
}

// removeTasks removes tasks of the passed in positions
func removeTasks(tasks []td.Task, positions []int) []td.Task {
	return slices.DeleteFunc(tasks, func(t td.Task) bool {
		return slices.Contains(positions, t.Position)
	})
}

func runRemove(cmd *cobra.Command, args []string) {
	wpName, wpFile, err := td.GetWorkingProject()
	if err != nil {
		log.Fatalf("unable to get current working project: %v\n", err)
	}

	tasks, err := td.ReadTasks(wpFile)
	if err != nil {
		log.Fatalf("unable to read tasks: %v\n", err)
	}

	if removeAllOpt {
		tasks = []td.Task{}
	} else {
		if len(args) < 1 {
			fmt.Println("Provide at least one task (number) to remove")
			return
		}
		positions, err := td.ArgsToPositions(args, len(tasks))
		if err != nil {
			fmt.Println(err)
			return
		}

		tasks = removeTasks(tasks, positions)
		td.OrderPositions(tasks)
	}

	err = td.WriteTasks(wpFile, tasks)
	if err != nil {
		log.Fatalf("unable to write tasks: %v\n", err)
	}

	td.PrintHeader(wpName)
	td.PrintTasks(tasks, td.ShowAll)
}
