package dashboard

type panelState int

const (
	waitingPanel panelState = iota
	runningPanel
	finishedPanel
)
