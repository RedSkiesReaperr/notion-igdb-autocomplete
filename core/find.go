package core

import (
	"fmt"

	"github.com/google/uuid"
)

func (c *Core) FindTasksByStatus(status TaskStatus) []Task {
	tasks := []Task{}

	for _, task := range c.Tasks {
		if task.Status == status {
			tasks = append(tasks, *task)
		}
	}

	return tasks
}

func (c *Core) FindTaskById(id string) (*Task, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid task Id: %v", id)
	}

	for _, task := range c.Tasks {
		if task.Id == parsedId {
			return task, nil
		}
	}

	return nil, fmt.Errorf("cannot find task with Id=%v", parsedId)
}
