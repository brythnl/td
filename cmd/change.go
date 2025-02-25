package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/brythnl/td/td"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// changeCmd represents the change command
var changeCmd = &cobra.Command{
	Use:     "change",
	Aliases: []string{"c"},
	Short:   "Change current project",
	Long:    `Change current project`,
	Args:    cobra.ExactArgs(1),
	Run:     runChange,
}

func init() {
	rootCmd.AddCommand(changeCmd)
}

func runChange(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatalln("Provide the project name to change to")
	}

	projectName := args[0]
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to detect home directory: %v\n", err)
	}

	projectFile := filepath.Join(home, ".td", "projects", projectName+".json")

	// Check if the project file exists, create it if it doesn't
	if _, err := os.Stat(projectFile); os.IsNotExist(err) {
		fmt.Printf("Project '%s' does not exist. Creating it...\n", projectName)
		emptyTasks := []td.Task{}
		if err := td.WriteTasks(projectFile, emptyTasks); err != nil {
			log.Fatalf("Failed to create project file: %v\n", err)
		}
	} else if err != nil {
		log.Fatalf("Error checking project file: %v\n", err)
	}

	// Persist the project name in the config file
	viper.Set("project", projectName)
	if err := viper.WriteConfig(); err != nil {
		log.Fatalf("Failed to write config file: %v\n", err)
	}

	fmt.Printf("Successfully changed project to '%s'\n", projectName)
}
