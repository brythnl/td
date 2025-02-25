package cmd

import (
	"log"

	"github.com/brythnl/td/td"

	"slices"

	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:     "move",
	Aliases: []string{"mv"},
	Short:   "Move a task to another position",
	Long:    `Move a task to another position`,
	Run:     runMove,
}

func init() {
	rootCmd.AddCommand(moveCmd)
}

func runMove(cmd *cobra.Command, args []string) {
	project := td.GetProjectFile()
	tasks, err := td.ReadTasks(project)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	if len(args) != 2 {
		log.Fatalln(
			"Invalid number of arguments. Please provide the task number to move and the target position.",
		)
	}

	positions := argsToPositions(args, len(tasks))
	currIdx, targetIdx := positions[0]-1, positions[1]-1
	taskToMove := tasks[currIdx]

	// Remove the task to move
	tasks = slices.Delete(tasks, currIdx, currIdx+1)
	// Insert the task to move at the target position
	tasks = append(
		tasks[:targetIdx],
		append([]td.Task{taskToMove}, tasks[targetIdx:]...)...)

	td.OrderPositions(tasks)

	err = td.WriteTasks(project, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	td.ShowTasks(tasks, td.ShowAll)
}
