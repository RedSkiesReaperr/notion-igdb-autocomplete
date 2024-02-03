package tui

type BackMsg struct{}

type SaveConfigMsg struct {
	NotionApiSecret string
	NotionPageId    string
	IgdbClientId    string
	IgdbSecret      string
	RefreshDelay    string
}
