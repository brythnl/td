package cmd

import (
	"log"

	"github.com/brythnl/td/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	dataFile := viper.GetString("datafile")
	tasks, err := todo.ReadTasks(dataFile)
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
	tasks = append(tasks[:currIdx], tasks[currIdx+1:]...)
	// Insert the task to move at the target position
	tasks = append(
		tasks[:targetIdx],
		append([]todo.Task{taskToMove}, tasks[targetIdx:]...)...)

	todo.OrderPositions(tasks)

	err = todo.WriteTasks(dataFile, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	showTasks(tasks, true)
}
