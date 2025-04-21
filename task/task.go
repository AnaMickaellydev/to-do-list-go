package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int
	Name     string
	Complete bool
}

func NewTask(id int, name string) Task {
	return Task{
		ID:       id,
		Name:     name,
		Complete: false,
	}
}

func (t Task) Print() {
	status := "❌"
	if t.Complete {
		status = "✅"
	}
	fmt.Printf("[%s] %d - %s/n", status, t.ID, t.Name)
}

func SaveTasksToFile(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadTasksFromFile(filename string) ([]Task, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {

			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
