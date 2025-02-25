package cmd

import (
	"log"
	"strconv"

	"github.com/brythnl/td/todo"
	"github.com/spf13/cobra"
)

var uncheckAllOpt bool

// uncheckCmd represents the uncheck command
var uncheckCmd = &cobra.Command{
	Use:     "uncheck",
	Aliases: []string{"ux"},
	Short:   "Uncheck a task",
	Long:    `Uncheck a task`,
	Run:     runUncheck,
}

func init() {
	rootCmd.AddCommand(uncheckCmd)

	uncheckCmd.Flags().BoolVarP(&uncheckAllOpt, "all", "a", false, "all tasks")
}

func runUncheck(cmd *cobra.Command, args []string) {
	project := todo.GetProjectFile()
	tasks, err := todo.ReadTasks(project)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	if uncheckAllOpt {
		for i := range tasks {
			tasks[i].Checked = false
		}
	} else {
		if len(args) < 1 {
			log.Fatalln("Provide at least one task (number) to uncheck")
		}

		for _, arg := range args {
			i, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatalln(arg, "is not a valid task number -", err)
			}
			if i < 1 || i > len(tasks) {
				log.Fatalln("Task", args[0], "is not available in the list")
			}
			tasks[i-1].Checked = false
		}
	}

	err = todo.WriteTasks(project, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	todo.ShowTasks(tasks, todo.ShowAll)
}
