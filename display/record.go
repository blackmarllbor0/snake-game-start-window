package display

import (
	"github.com/gdamore/tcell/v2"
	tview "github.com/rivo/tview"
)

type RecordDisplay struct {
	label       *tview.TextView
	bestPlayers *tview.List
	block       *tview.Flex
}

// init инициализирует блок
func (r *RecordDisplay) init() *tview.Flex {
	r.block = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(r.createLabel(), 0, 1, false)
	return r.block
}

// createLabel создает заголовок
func (r *RecordDisplay) createLabel() *tview.TextView {
	r.label = tview.NewTextView().
		SetText("RECORDS").
		SetTextAlign(tview.AlignCenter).
		SetTextStyle(tcell.StyleDefault.
			Background(cw).
			Bold(true))
	r.label.SetTextColor(tview.AlignCenter)
	r.label.SetBackgroundColor(cb)
	return r.label
}
