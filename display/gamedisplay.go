package display

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GameDisplay struct {
	label     *tview.TextView
	inputName *tview.InputField
	block     *tview.Flex
}

// init инициализирует блок
func (g *GameDisplay) init() *tview.Flex {
	var (
		save  = g.createAction("Ent", "save")
		help  = g.createAction("h", "help")
		start = g.createAction("s", "start")
		quit  = g.createAction("q", "quit")
	)

	// это для выравнивания инпута
	var flex = tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(g.createInput(), 20, 0, true).
		AddItem(tview.NewBox(), 0, 1, false)

	g.block = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(g.createLabel(), 5, 2, false).
		AddItem(flex, 0, 1, true).
		AddItem(save, 0, 1, false).
		AddItem(help, 0, 1, false).
		AddItem(start, 0, 1, false).
		AddItem(quit, 0, 1, false)
	return g.block
}

// createInput создает поле для ввода имени
func (g *GameDisplay) createInput() *tview.InputField {
	g.inputName = tview.NewInputField().
		SetLabel("Enter name: ").
		SetLabelColor(cw).
		SetFieldBackgroundColor(cb).
		SetFieldWidth(15)
	return g.inputName
}

// createLabel создает заголовок
func (g *GameDisplay) createLabel() *tview.TextView {
	g.label = tview.NewTextView().
		SetText("GAME").
		SetTextStyle(tcell.StyleDefault.
			Background(cw).
			Bold(true))
	g.label.SetTextAlign(tview.AlignCenter)
	g.label.SetBackgroundColor(cb)
	return g.label
}

// createAction создает поля для действий
func (g *GameDisplay) createAction(key string, cmd string) *tview.TextView {
	return tview.NewTextView().
		SetText(fmt.Sprintf("Press (%s) to %s\n", key, cmd)).
		SetTextColor(tcell.ColorWhite).
		SetTextAlign(tview.AlignCenter)
}
