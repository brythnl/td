package cmd

import (
	"fmt"
	"log"

	"github.com/brythnl/td/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	listAllOpt     bool
	listCheckedOpt bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List the current tasks",
	Long:    `List the current tasks`,
	Run:     runList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&listAllOpt, "all", "a", false, "all tasks")
	listCmd.Flags().BoolVarP(&listCheckedOpt, "checked", "x", false, "checked tasks")
}

// showTasks prints the tasks in the given slice.
//
// If showAll is true, all tasks are shown (checked and unchecked).
func showTasks(tasks []todo.Task, showAll bool) {
	fmt.Println()
	if len(tasks) == 0 {
		fmt.Println("All done!")
		return
	}

	for _, t := range tasks {
		// Show only unchecked tasks by default
		if showAll || listAllOpt || t.Checked == listCheckedOpt {
			fmt.Print(t.Prefix(), t.Text, "\n\n")
		}
	}
}

func runList(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	tasks, err := todo.ReadTasks(dataFile)
	if err != nil {
		log.Printf("%v\n", err)
	}

	showTasks(tasks, false)
}
