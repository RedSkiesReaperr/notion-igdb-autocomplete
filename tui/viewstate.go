package tui

type ViewState int

const (
	MainView ViewState = iota
	DashboardView
	ConfigurationView
	VerifyConfigurationView
)
