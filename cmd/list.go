package cmd

import (
	"fmt"
	"log"

	"github.com/brythnl/td/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	allOpt     bool
	checkedOpt bool
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List the current tasks",
	Long:    `List the current tasks`,
	Run:     runList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "Show all")
	listCmd.Flags().BoolVarP(&checkedOpt, "checked", "x", false, "Show checked")
}

func runList(cmd *cobra.Command, args []string) {
	dataFile := viper.GetString("datafile")
	tasks, err := todo.ReadTasks(dataFile)
	if err != nil {
		log.Printf("%v\n", err)
	}

	for _, t := range tasks {
		// Show only unchecked tasks by default
		if allOpt || t.Checked == checkedOpt {
			fmt.Println(t.Prefix(), t.Text)
		}
	}
}
