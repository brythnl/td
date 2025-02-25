package cmd

import (
	"log"
	"slices"

	"github.com/brythnl/td/todo"

	"github.com/spf13/cobra"
)

var removeAllOpt bool

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a task",
	Long:    `Remove a task`,
	Run:     runRemove,
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolVarP(&removeAllOpt, "all", "a", false, "all tasks")
}

// removeTasks removes tasks of the passed in positions
func removeTasks(tasks []todo.Task, positions []int) []todo.Task {
	return slices.DeleteFunc(tasks, func(t todo.Task) bool {
		return slices.Contains(positions, t.Position)
	})
}

func runRemove(cmd *cobra.Command, args []string) {
	project := todo.GetProjectFile()
	tasks, err := todo.ReadTasks(project)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	if removeAllOpt {
		tasks = []todo.Task{}
	} else {
		if len(args) < 1 {
			log.Fatalln("Provide at least one task (number) to remove")
		}

		tasks = removeTasks(tasks, argsToPositions(args, len(tasks)))
		todo.OrderPositions(tasks)
	}

	err = todo.WriteTasks(project, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	todo.ShowTasks(tasks, todo.ShowAll)
}
