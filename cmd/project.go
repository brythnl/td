package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/brythnl/td/td"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var projectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Project-related commands",
}

var listProjectsCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List the current existing projects",
	Run:     runListProjects,
}

var changeProjectCmd = &cobra.Command{
	Use:     "change",
	Aliases: []string{"c"},
	Short:   "Change current project",
	Args:    cobra.ExactArgs(1),
	Run:     runChangeProject,
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(changeProjectCmd, listProjectsCmd)
}

func runListProjects(cmd *cobra.Command, args []string) {
	// Get all existing projects
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to detect home directory: %v\n", err)
	}
	projects, err := os.ReadDir(filepath.Join(home, ".td", "projects"))
	if err != nil {
		log.Fatalf("Unable to read projects directory: %v\n", err)
	}
	// Get working project name
	wpName, _, err := td.GetProject()
	if err != nil {
		log.Fatalf("Unable to get current working project: %v\n", err)
	}

	fmt.Println("========")
	fmt.Println("Projects")
	fmt.Println("========")

	for _, p := range projects {
		pName := strings.TrimSuffix(p.Name(), ".json")
		fmt.Print("• " + pName)
		if wpName == pName {
			fmt.Print(" 🡄")
		}
		fmt.Println()
	}
}

func runChangeProject(cmd *cobra.Command, args []string) {
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
		log.Fatalf("Failed to set project: %v\n", err)
	}

	fmt.Printf("Successfully changed project to '%s'\n", projectName)
}
