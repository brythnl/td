package cmd

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/brythnl/td/td"
	"github.com/spf13/cobra"
)

var checkAllOpt bool

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"x"},
	Short:   "Check a task to mark it as done",
	Run:     runCheck,
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().BoolVarP(&checkAllOpt, "all", "a", false, "all tasks")
}

func runCheck(cmd *cobra.Command, args []string) {
	wpName, wpFile, err := td.GetWorkingProject()
	if err != nil {
		log.Fatalf("unable to get current working project: %v\n", err)
	}
	tasks, err := td.ReadTasks(wpFile)
	if err != nil {
		log.Fatalf("unable to read tasks: %v\n", err)
	}

	if checkAllOpt {
		for i := range tasks {
			tasks[i].Checked = true
		}
	} else {
		if len(args) < 1 {
			fmt.Println("Provide at least one task (number) to check")
			return
		}

		for _, arg := range args {
			i, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(arg, "is not a valid task number -", err)
				return
			}
			if i < 1 || i > len(tasks) {
				fmt.Println("Task", args[0], "is not available in the list")
				return
			}

			checkedTask := tasks[i-1]
			checkedTask.Checked = true

			// Move checked task to end of list
			// BUG: when the user checks multiple tasks and it includes the last task,
			// each task to be checked is moved to the end of the list, thus replacing itself
			tasks = slices.Delete(tasks, i-1, i)
			tasks = append(tasks, checkedTask)
			td.OrderPositions(tasks)
		}
	}

	if err = td.WriteTasks(wpFile, tasks); err != nil {
		log.Fatalf("unable to write tasks: %v\n", err)
	}

	td.PrintHeader(wpName)
	td.PrintTasks(tasks, td.ShowAll)
}
