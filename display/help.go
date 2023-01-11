package display

import (
	"github.com/rivo/tview"
)

type HelpModal struct {
	modal *tview.Modal
}

func (h *HelpModal) init() *tview.Modal {
	h.modal = tview.NewModal().
		SetBackgroundColor(cb)
	h.modal.SetBorderColor(cw)
	h.modal.SetText("I'm sorry, but it's nothing\n;(\nPress (q) to back")
	return h.modal
}
