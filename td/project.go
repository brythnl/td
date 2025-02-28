package td

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// GetProject gets path of the project file and the project name
func GetProject() (string, string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	project := viper.GetString("project")
	return project, filepath.Join(home, ".td", "projects", project+".json"), nil
}
