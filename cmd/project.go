package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/brythnl/td/td"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Project-related commands",
}

var listProjectsCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List existing projects",
	Run:     runListProjects,
}

var switchProjectCmd = &cobra.Command{
	Use:     "switch",
	Aliases: []string{"s"},
	Short:   "Switch current working project",
	Run:     runSwitchProject,
}

var renameProjectCmd = &cobra.Command{
	Use:     "rename",
	Aliases: []string{"r"},
	Short:   "Rename current working project",
	Run:     runRenameProject,
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(
		switchProjectCmd,
		listProjectsCmd,
		renameProjectCmd,
	)
}

func runListProjects(cmd *cobra.Command, args []string) {
	// Get all existing projects
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable to detect home directory: %v\n", err)
	}
	projects, err := os.ReadDir(filepath.Join(home, ".td", "projects"))
	if err != nil {
		log.Fatalf("unable to read projects directory: %v\n", err)
	}

	// Get working project name
	wpName, _, err := td.GetWorkingProject()
	if err != nil {
		log.Fatalf("unable to get current working project: %v\n", err)
	}

	fmt.Println("========")
	fmt.Println("Projects")
	fmt.Println("========")

	for _, p := range projects {
		pName := strings.TrimSuffix(p.Name(), ".json")
		fmt.Print("• ", pName)
		if wpName == pName {
			fmt.Print(" 🡄")
		}
		fmt.Println()
	}
}

func runSwitchProject(cmd *cobra.Command, args []string) {
	// Get target project name
	if len(args) != 1 {
		fmt.Println("Provide ONE project name to switch to")
	}
	tpName := args[0]

	// Check if the project file exists, create it if it doesn't
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable to detect home directory: %v\n", err)
	}

	tpFile := filepath.Join(home, ".td", "projects", tpName+".json")

	if _, err := os.Stat(tpFile); os.IsNotExist(err) {
		fmt.Printf("Project '%s' does not exist. Creating it...\n", tpName)
		emptyTasks := []td.Task{}
		if err := td.WriteTasks(tpFile, emptyTasks); err != nil {
			log.Fatalf("unable to create project file: %v\n", err)
		}
	} else if err != nil {
		log.Fatalf("error checking project file: %v\n", err)
	}

	// Set as current working project
	if err = td.SetWorkingProject(tpName); err != nil {
		log.Fatalf("unable to set working project: %v\n", err)
	}

	tasks, err := td.ReadTasks(tpFile)
	if err != nil {
		log.Fatalf("unable to read tasks: %v\n", err)
	}

	td.PrintHeader(tpName)
	td.PrintTasks(tasks, td.ShowUnchecked)
}

func runRenameProject(cmd *cobra.Command, args []string) {
	// Get new project name
	if len(args) != 1 {
		fmt.Println("Provide ONE new name for the current working project")
		return
	}
	npName := args[0]

	// Get working project file
	_, wpFile, err := td.GetWorkingProject()
	if err != nil {
		log.Fatalf("unable to get current working project: %v\n", err)
	}

	// Check if there is an existing project with the same name
	if _, err = os.ReadFile(filepath.Join(filepath.Dir(wpFile), npName+".json")); err == nil {
		fmt.Printf("Project '%s' already exists. Choose another name\n", npName)
		return
	}

	// Rename project file
	err = os.Rename(wpFile, filepath.Join(filepath.Dir(wpFile), npName+".json"))
	if err != nil {
		log.Fatalf("unable to rename project file: %v\n", err)
	}

	// Rename project in config
	if err = td.SetWorkingProject(npName); err != nil {
		log.Fatalf("unable to set working project: %v\n", err)
	}

	fmt.Println("Project renamed to '", npName, "'")
}
