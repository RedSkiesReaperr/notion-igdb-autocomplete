package tui

type BackMsg struct {
	Width  int
	Height int
}

type SaveConfigMsg struct {
	NotionApiSecret string
	NotionPageId    string
	IgdbClientId    string
	IgdbSecret      string
	RefreshDelay    string
}
