package td

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Task struct {
	Text     string
	Position int
	Checked  bool
}

// Prefix returns prefix of each task line.
//
// e.g.: [ ] 1 >
func (t *Task) Prefix() string {
	checkbox := "[ ]"
	if t.Checked {
		checkbox = "[x]"
	}
	return checkbox + " " + strconv.Itoa(t.Position) + " > "
}

// OrderPositions reorders the Position field of the tasks.
func OrderPositions(tasks []Task) {
	for i := range tasks {
		tasks[i].Position = i + 1
	}
}

// GetProjectFile gets path of the project file
func GetProjectFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to detect home directory: %v\n", err)
	}
	project := viper.GetString("project")
	return filepath.Join(home, ".td", "projects", project+".json")
}

// WriteTasks writes tasks data to JSON file.
func WriteTasks(filename string, tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return err
	}

	return nil
}

// ReadTasks reads tasks data from JSON file.
func ReadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

type ShowOption int

const (
	ShowAll ShowOption = iota
	ShowChecked
	ShowUnchecked
)

// showTasks prints the tasks in the given slice.
func ShowTasks(tasks []Task, opt ShowOption) {
	header := "Project: " + viper.GetString("project")
	sep := strings.Repeat("=", len(header))
	fmt.Println(sep)
	fmt.Println(header)
	fmt.Println(sep)

	if len(tasks) == 0 {
		fmt.Println("All done!")
		return
	}

	for _, t := range tasks {
		switch opt {
		case ShowAll:
			fmt.Print(t.Prefix(), t.Text, "\n\n")
		case ShowChecked:
			if t.Checked {
				fmt.Print(t.Prefix(), t.Text, "\n\n")
			}
		case ShowUnchecked:
			if !t.Checked {
				fmt.Print(t.Prefix(), t.Text, "\n\n")
			}
		}
	}
}
