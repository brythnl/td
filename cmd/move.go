package cmd

import (
	"fmt"
	"log"
	"slices"

	"github.com/brythnl/td/td"
	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:     "move",
	Aliases: []string{"mv"},
	Short:   "Move a task to another position",
	Run:     runMove,
}

func init() {
	rootCmd.AddCommand(moveCmd)
}

func runMove(cmd *cobra.Command, args []string) {
	wpName, wpFile, err := td.GetWorkingProject()
	if err != nil {
		log.Fatalf("unable to get current working project: %v\n", err)
	}
	tasks, err := td.ReadTasks(wpFile)
	if err != nil {
		log.Fatalf("unable to read tasks: %v\n", err)
	}

	if len(args) != 2 {
		fmt.Println(
			"Invalid number of arguments. Please provide the task number to move and the target position.",
		)
		return
	}

	positions, err := td.ArgsToPositions(args, len(tasks))
	if err != nil {
		fmt.Println(err)
		return
	}

	currIdx, targetIdx := positions[0]-1, positions[1]-1
	taskToMove := tasks[currIdx]

	// Remove the task to move
	tasks = slices.Delete(tasks, currIdx, currIdx+1)
	// Insert the task to move at the target position
	tasks = append(
		tasks[:targetIdx],
		append([]td.Task{taskToMove}, tasks[targetIdx:]...)...)

	td.OrderPositions(tasks)

	if err = td.WriteTasks(wpFile, tasks); err != nil {
		log.Fatalf("unable to write tasks: %v\n", err)
	}

	td.PrintHeader(wpName)
	td.PrintTasks(tasks, td.ShowAll)
}
