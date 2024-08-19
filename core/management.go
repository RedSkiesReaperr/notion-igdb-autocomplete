package core

import (
	"fmt"
	"log"
)

func (c *Core) registerTask(task *Task) error {
	if c.isRegisteredTask(task) {
		return fmt.Errorf("already exists: %v", task)
	}

	c.Tasks = append(c.Tasks, task)
	log.Printf("Registered task: %v\n", task)

	return nil
}

func (c *Core) isRegisteredTask(task *Task) bool {
	for _, t := range c.Tasks {
		if t.NotionId == task.NotionId && t.Type == task.Type && (task.Status == WaitingTask || task.Status == RunningTask) {
			return true
		}
	}

	return false
}
