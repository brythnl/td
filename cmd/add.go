package cmd

import (
	"fmt"
	"log"

	"github.com/brythnl/td/td"
	"github.com/spf13/cobra"
)

var addPositionOpt int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new task",
	Run:     func(cmd *cobra.Command, args []string) { add(args) },
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&addPositionOpt, "position", "p", -1, "position of the new task")
}

func add(args []string) {
	wpName, wpFile, err := td.GetWorkingProject()
	if err != nil {
		log.Fatalf("unable to get current working project: %v\n", err)
	}
	tasks, err := td.ReadTasks(wpFile)
	if err != nil {
		log.Fatalf("unable to read tasks: %v\n", err)
	}

	if addPositionOpt == -1 {
		for _, t := range args {
			task := td.Task{Text: t}
			tasks = append(tasks, task)
		}
	} else {
		if len(args) > 1 {
			fmt.Println("Too many arguments. When using the -p option, only one task can be added")
			return
		}

		targetIdx := addPositionOpt - 1
		task := td.Task{Text: args[0]}
		// Insert the task to move at the target position
		tasks = append(
			tasks[:targetIdx],
			append([]td.Task{task}, tasks[targetIdx:]...)...)
	}

	td.OrderPositions(tasks)

	if err := td.WriteTasks(wpFile, tasks); err != nil {
		log.Printf("unable to write tasks: %v\n", err)
	}

	td.PrintHeader(wpName)
	td.PrintTasks(tasks, td.ShowUnchecked)
}
