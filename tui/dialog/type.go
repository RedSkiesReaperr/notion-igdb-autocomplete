package dialog

type DialogType int

const (
	SuccessDialog DialogType = iota
	ErrorDialog
	InfoDialog
)
