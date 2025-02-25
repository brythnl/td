package cmd

import (
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
	Long:    `Remove a task`,
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
	project := td.GetProjectFile()
	tasks, err := td.ReadTasks(project)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	if removeAllOpt {
		tasks = []td.Task{}
	} else {
		if len(args) < 1 {
			log.Fatalln("Provide at least one task (number) to remove")
		}

		tasks = removeTasks(tasks, argsToPositions(args, len(tasks)))
		td.OrderPositions(tasks)
	}

	err = td.WriteTasks(project, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	td.ShowTasks(tasks, td.ShowAll)
}
