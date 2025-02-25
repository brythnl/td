package cmd

import (
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
	Long:    `Add a new task`,
	Run:     func(cmd *cobra.Command, args []string) { add(args) },
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&addPositionOpt, "position", "p", -1, "position of the new task")
}

func add(args []string) {
	project := td.GetProjectFile()
	tasks, err := td.ReadTasks(project)
	if err != nil {
		log.Printf("%v\n", err)
	}

	if addPositionOpt == -1 {
		for _, t := range args {
			task := td.Task{Text: t}
			tasks = append(tasks, task)
		}
	} else {
		if len(args) > 1 {
			log.Fatalln("Too many arguments. When using the -p option, only one task can be added.")
		}

		targetIdx := addPositionOpt - 1
		task := td.Task{Text: args[0]}
		// Insert the task to move at the target position
		tasks = append(
			tasks[:targetIdx],
			append([]td.Task{task}, tasks[targetIdx:]...)...)
	}

	td.OrderPositions(tasks)

	if err := td.WriteTasks(project, tasks); err != nil {
		log.Printf("%v\n", err)
	}

	td.ShowTasks(tasks, td.ShowUnchecked)
}
