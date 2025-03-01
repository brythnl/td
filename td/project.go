package td

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// GetWorkingProject gets the name and file path of the current working project
func GetWorkingProject() (string, string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	project := viper.GetString("project")
	return project, filepath.Join(home, ".td", "projects", project+".json"), nil
}

func SetWorkingProject(projectName string) error {
	viper.Set("project", projectName)
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("Failed to set project: %v\n", err)
	}
	return nil
}

// PrintHeader prints the name of the given project as a header
func PrintHeader(projectName string) {
	header := "Project: " + projectName
	sep := strings.Repeat("=", len(header))
	fmt.Println(sep)
	fmt.Println(header)
	fmt.Println(sep)
}
