package td

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

// printTasks prints the tasks in the given tasks slice.
func PrintTasks(tasks []Task, opt ShowOption) {
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

// argsToPositions converts passed arguments (strings) into positions (integers).
func ArgsToPositions(args []string, tasksCount int) ([]int, error) {
	positions := make([]int, 0, len(args))
	for _, arg := range args {
		p, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf(arg + " is not a valid task number")
		}
		if p < 1 || p > tasksCount {
			return nil, fmt.Errorf("Task " + arg + " is not available in the list")
		}
		positions = append(positions, p)
	}

	return positions, nil
}
