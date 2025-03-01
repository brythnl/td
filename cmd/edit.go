package cmd

import (
	"fmt"
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
	Run:     runEdit,
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func runEdit(cmd *cobra.Command, args []string) {
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
			"Invalid number of arguments. Please provide a task number and a new description.",
		)
		return
	}
	p, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(
			"Invalid task number. Please provide a valid task number for the first argument.",
		)
		return
	}
	if p < 1 || p > len(tasks) {
		fmt.Println("Task", p, "is not available in the list")
		return
	}

	tasks[p-1].Text = args[1]

	if err = td.WriteTasks(wpFile, tasks); err != nil {
		log.Fatalf("unable to write tasks: %v\n", err)
	}

	td.PrintHeader(wpName)
	td.PrintTasks(tasks, td.ShowUnchecked)
}
