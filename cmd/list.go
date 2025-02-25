package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/brythnl/td/todo"

	"github.com/spf13/cobra"
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

func runList(cmd *cobra.Command, args []string) {
	option := todo.ShowUnchecked
	if listAllOpt {
		option = todo.ShowAll
	} else if listCheckedOpt {
		option = todo.ShowChecked
	}

	if len(args) == 0 {
		project = todo.GetProjectFile()
		tasks, err := todo.ReadTasks(project)
		if err != nil {
			log.Printf("%v\n", err)
		}

		todo.ShowTasks(tasks, option)
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to detect home directory: %v\n", err)
		return
	}

	for _, arg := range args {
		projectName := arg
		project = filepath.Join(home, ".td", "projects", projectName+".json")
		tasks, err := todo.ReadTasks(project)
		if err != nil {
			log.Printf("%v\n", err)
		}

		todo.ShowTasks(tasks, option)
		return
	}
}
