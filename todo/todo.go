package todo

import (
	"encoding/json"
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
