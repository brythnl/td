package cmd

import (
	"log"

	"github.com/brythnl/td/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task`,
	Run:   runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	tasks, err := todo.ReadTasks(dataFile)
	if err != nil {
		log.Printf("%v\n", err)
	}

	for _, t := range args {
		task := todo.Task{Text: t}
		tasks = append(tasks, task)
	}
	todo.OrderPositions(tasks)

	if err := todo.WriteTasks(dataFile, tasks); err != nil {
		log.Printf("%v\n", err)
	}

	showTasks(tasks, false)
}
