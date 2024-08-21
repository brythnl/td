package cmd

import (
	"log"
	"strconv"

	"github.com/brythnl/td/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var checkCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"x"},
	Short:   "Check a task (number) to mark it as done",
	Long:    `Check a task (number) to mark it as done`,
	Run:     runCheck,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func runCheck(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	tasks, err := todo.ReadTasks(dataFile)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
	}

	if len(args) < 1 {
		log.Fatalln("Invalid number of arguments")
	}
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalln(arg, "is not a valid task number -", err)
		}
		if i < 1 || i > len(tasks) {
			log.Fatalln("Task", args[0], "is not available in the list")
		}

		tasks[i-1].Checked = true
	}

	err = todo.WriteTasks(dataFile, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	showTasks(tasks, true)
}
