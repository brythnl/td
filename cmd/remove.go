package cmd

import (
	"log"
	"slices"
	"strconv"

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

func RemoveTasks(tasks []todo.Task, positions []int) []todo.Task {
	return slices.DeleteFunc(tasks, func(t todo.Task) bool {
		for _, p := range positions {
			if t.Position == p {
				return true
			}
		}
		return false
	})
}

// argsToPositions converts passed arguments (strings) into positions (integers).
func argsToPositions(args []string, tasksCount int) []int {
	if len(args) < 1 {
		log.Fatalln("Provide at least one argument (task number)")
	}

	positions := make([]int, len(args))
	for _, arg := range args {
		p, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalln(arg, "is not a valid task number -", err)
		}
		if p < 1 || p > tasksCount {
			log.Fatalln("Task", arg, "is not available in the list")
		}
		positions = append(positions, p)
	}

	return positions
}

func runRemove(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	tasks, err := todo.ReadTasks(dataFile)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	tasks = RemoveTasks(tasks, argsToPositions(args, len(tasks)))
	todo.OrderPositions(&tasks)

	err = todo.WriteTasks(dataFile, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}
}
