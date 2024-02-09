package dashboard

import "github.com/charmbracelet/bubbles/table"

type Table struct {
	model   table.Model
	columns []table.Column
	rows    []table.Row
}
