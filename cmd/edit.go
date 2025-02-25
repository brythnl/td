package cmd

import (
	"log"
	"strconv"

	"github.com/brythnl/td/td"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e", "ed"},
	Short:   "Edit a task",
	Long:    `Edit a task`,
	Run:     runEdit,
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func runEdit(cmd *cobra.Command, args []string) {
	project := td.GetProjectFile()
	tasks, err := td.ReadTasks(project)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	if len(args) != 2 {
		log.Fatalln(
			"Invalid number of arguments. Please provide a task number and a new description.",
		)
	}
	p, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(
			"Invalid task number. Please provide a valid task number for the first argument.",
		)
	}
	if p < 1 || p > len(tasks) {
		log.Fatalln("Task", p, "is not available in the list")
	}

	tasks[p-1].Text = args[1]

	err = td.WriteTasks(project, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	td.ShowTasks(tasks, td.ShowUnchecked)
}
