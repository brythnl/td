package cmd

import (
	"log"
	"strconv"

	"github.com/brythnl/td/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checkAllOpt bool

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"x"},
	Short:   "Check a task to mark it as done",
	Long:    `Check a task to mark it as done`,
	Run:     runCheck,
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().BoolVarP(&checkAllOpt, "all", "a", false, "all tasks")
}

func runCheck(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	tasks, err := todo.ReadTasks(dataFile)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	if checkAllOpt {
		for i := range tasks {
			tasks[i].Checked = true
		}
	} else {
		if len(args) < 1 {
			log.Fatalln("Provide at least one task (number) to check")
		}

		for _, arg := range args {
			i, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatalln(arg, "is not a valid task number -", err)
			}
			if i < 1 || i > len(tasks) {
				log.Fatalln("Task", args[0], "is not available in the list")
			}

			checkedTask := tasks[i-1]
			checkedTask.Checked = true

			// Move checked task to end of list
			tasks = append(tasks[:i-1], tasks[i:]...)
			tasks = append(tasks, checkedTask)
			todo.OrderPositions(tasks)
		}
	}

	err = todo.WriteTasks(dataFile, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	showTasks(tasks, true)
}
