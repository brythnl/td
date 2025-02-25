package cmd

import (
	"log"

	"github.com/brythnl/td/todo"

	"github.com/spf13/cobra"
)

var addPositionOpt int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new task",
	Long:    `Add a new task`,
	Run:     func(cmd *cobra.Command, args []string) { add(args) },
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&addPositionOpt, "position", "p", -1, "position of the new task")
}

func add(args []string) {
	project := todo.GetProjectFile()
	tasks, err := todo.ReadTasks(project)
	if err != nil {
		log.Printf("%v\n", err)
	}

	if addPositionOpt == -1 {
		for _, t := range args {
			task := todo.Task{Text: t}
			tasks = append(tasks, task)
		}
	} else {
		if len(args) > 1 {
			log.Fatalln("Too many arguments. When using the -p option, only one task can be added.")
		}

		targetIdx := addPositionOpt - 1
		task := todo.Task{Text: args[0]}
		// Insert the task to move at the target position
		tasks = append(
			tasks[:targetIdx],
			append([]todo.Task{task}, tasks[targetIdx:]...)...)
	}

	todo.OrderPositions(tasks)

	if err := todo.WriteTasks(project, tasks); err != nil {
		log.Printf("%v\n", err)
	}

	todo.ShowTasks(tasks, todo.ShowUnchecked)
}
