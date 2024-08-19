package core

type TaskStatus int

const (
	WaitingTask TaskStatus = iota
	RunningTask
	FinishedTask
)

func TaskStatusString(status TaskStatus) string {
	switch status {
	case WaitingTask:
		return "Waiting"
	case RunningTask:
		return "Running"
	case FinishedTask:
		return "Finished"
	default:
		return "Unknown"
	}
}
