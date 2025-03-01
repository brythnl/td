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
		wpName, wpFile, err := td.GetWorkingProject()
		if err != nil {
			log.Fatalf("unable to get current working project: %v\n", err)
		}
		tasks, err := td.ReadTasks(wpFile)
		if err != nil {
			log.Fatalf("unable to read tasks: %v\n", err)
		}

		td.PrintHeader(wpName)
		td.PrintTasks(tasks, option)
		return
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable to detect home directory: %v\n", err)
	}

	for _, arg := range args {
		pName := arg
		pFile := filepath.Join(home, ".td", "projects", pName+".json")
		// TODO: check if pFile exists
		tasks, err := td.ReadTasks(pFile)
		if err != nil {
			log.Fatalf("unable to read tasks: %v\n", err)
		}

		td.PrintHeader(pName)
		td.PrintTasks(tasks, option)
	}
}
