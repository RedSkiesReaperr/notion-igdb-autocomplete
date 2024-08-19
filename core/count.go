package core

func (c *Core) CountTasksWithStatus(status TaskStatus) int {
	return len(c.FindTasksByStatus(status))
}
