package display

import (
	"snake/customerror"
	"snake/users"

	"github.com/gdamore/tcell/v2"
	tview "github.com/rivo/tview"
)

type Display struct {
	App      *tview.Application
	Pages    *tview.Pages
	Players  []users.Player
	MainFlex *tview.Flex
	GameFlex *tview.Flex
}

const (
	menu    = "Menu"
	addUser = "Add user"
)

func (d *Display) Displaying() {
	d.Players = users.GetPlayers() // получаем игроков
	d.App = tview.NewApplication() // Графическое приложение

	d.gameFlex()

	d.MainFlex = tview.NewFlex().
		SetDirection(tview.FlexColumnCSS).
		AddItem(d.GameFlex, 0, 1, false)

	if err := d.App.SetRoot(d.MainFlex, true).Run(); err != nil {
		panic(customerror.ErrorString(customerror.CustomError{
			Message: "Приложение упало при запуске",
			Error:   err,
		}))
	}
}

func (d *Display) gameFlex() {
	var (
		game  = createTextView("Игра")
		name  = createTextView("Введите имя:")
		help  = createTextView("Нажмите (h) для справки")
		start = createTextView("Нажмите (s) для старта")
		quit  = createTextView("Нажмите (q) для выхода")
	)

	d.GameFlex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(game, 4, 2, false).
		AddItem(name, 2, 1, true).
		AddItem(help, 2, 1, false).
		AddItem(start, 2, 1, false).
		AddItem(quit, 0, 1, false)
}

func createTextView(text string) *tview.TextView {
	return tview.NewTextView().
		SetText(text).
		SetTextColor(tcell.ColorWhite).
		SetTextAlign(tview.AlignCenter)
}
