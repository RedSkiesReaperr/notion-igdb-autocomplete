package core

type TaskType int

const (
	GameInfosTask TaskType = iota
	TimeToBeatTask
)

func TaskTypeString(t TaskType) string {
	switch t {
	case GameInfosTask:
		return "GameInfos"
	case TimeToBeatTask:
		return "TimeToBeat"
	default:
		return "Unknown"
	}
}
