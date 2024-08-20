package cmd

import (
	"log"
	"strconv"

	"github.com/brythnl/td/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var uncheckCmd = &cobra.Command{
	Use:     "uncheck",
	Aliases: []string{"ux"},
	Short:   "Uncheck a task (number)",
	Long:    `Uncheck a task (number)`,
	Run:     runUncheck,
}

func init() {
	rootCmd.AddCommand(uncheckCmd)
}

func runUncheck(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	tasks, err := todo.ReadTasks(dataFile)
	if err != nil {
		log.Fatalf("Read tasks error: %v\n", err)
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
	todo.OrderPositions(&tasks)

	err = todo.WriteTasks(dataFile, tasks)
	if err != nil {
		log.Fatalf("Write tasks error: %v\n", err)
	}

	showTasks(tasks, true)
}
