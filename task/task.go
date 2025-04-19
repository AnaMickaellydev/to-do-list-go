package task

import "fmt"

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
