package cmd

import (
	"log"
	"strconv"

	"github.com/brythnl/td/td"

	"slices"

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
	projectName, projectFile, err := td.GetProject()
	if err != nil {
		log.Fatalf("Unable to get current working project: %v\n", err)
	}
	tasks, err := td.ReadTasks(projectFile)
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
			// BUG: when the user checks multiple tasks and it includes the last task,
			// each task to be checked is moved to the end of the list, thus replacing itself
			tasks = slices.Delete(tasks, i-1, i)
			tasks = append(tasks, checkedTask)
			td.OrderPositions(tasks)
		}
	}

	err = td.WriteTasks(projectFile, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	td.ShowTasks(tasks, td.ShowAll, projectName)
}
