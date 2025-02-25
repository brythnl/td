package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/brythnl/td/td"

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
	option := td.ShowUnchecked
	if listAllOpt {
		option = td.ShowAll
	} else if listCheckedOpt {
		option = td.ShowChecked
	}

	if len(args) == 0 {
		projectName, projectFile := td.GetProject()
		tasks, err := td.ReadTasks(projectFile)
		if err != nil {
			log.Fatalf("%v\n", err)
		}

		td.ShowTasks(tasks, option, projectName)
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
		tasks, err := td.ReadTasks(project)
		if err != nil {
			log.Fatalf("%v\n", err)
		}

		td.ShowTasks(tasks, option, projectName)
		return
	}
}
