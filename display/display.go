package display

import (
	"snake/customerror"
	"snake/users"

	"github.com/gdamore/tcell/v2"
	tview "github.com/rivo/tview"
)

type Display struct {
	App        *tview.Application
	Pages      *tview.Pages
	Players    []users.Player
	MainFlex   *tview.Flex
	GameFlex   *tview.Flex
	RecordFlex *tview.Flex
	PlayerName string
}

const (
	menu    = "Menu"
	addUser = "Add user"
)

func (d *Display) Displaying() {
	d.Players = users.GetPlayers() // получаем игроков
	d.App = tview.NewApplication() // Графическое приложение
	d.Pages = tview.NewPages()     // init pages

	// d.Pages.AddPage("Menu", d.MainFlex, false, false)

	d.gameFlex()   // init game flex
	d.recordFlex() // init record flex

	a := tview.NewBox().
		SetBackgroundColor(tcell.ColorWhite)

	d.MainFlex = tview.NewFlex().
		SetDirection(tview.FlexColumn). // делим жкран на две части
		// и добовляем элементы:
		AddItem(d.GameFlex, 0, 2, true).
		AddItem(a, 1, 0, false).
		AddItem(d.RecordFlex, 0, 1, false)

	// запуск приложения
	if err := d.App.SetRoot(d.MainFlex, true).Run(); err != nil {
		panic(customerror.ErrorString(customerror.CustomError{
			Message: "Приложение упало при запуске",
			Error:   err,
		}))
	}
}

func (d *Display) gameFlex() {
	var (
		game = createTextView("GAME").
			SetTextStyle(tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Bold(true))
		name  = playerEnterName()
		save  = createTextView("Press (Ent) to save")
		help  = createTextView("Press (h) to help")
		start = createTextView("Press (s) to start")
		quit  = createTextView("Press (q) to quit")
	)

	d.GameFlex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(game, 4, 2, false).
		AddItem(name, 2, 1, true).
		AddItem(save, 2, 1, false).
		AddItem(help, 2, 1, false).
		AddItem(start, 2, 1, false).
		AddItem(quit, 0, 1, false)
}

func playerEnterName() *tview.Flex {
	name := tview.NewInputField().
		SetLabel("Enter name: ").
		SetLabelColor(tcell.ColorWhite).
		SetFieldBackgroundColor(tcell.ColorBlack)

	return tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(name, 0, 4, true).
		AddItem(tview.NewBox(), 0, 1, false)
}

func (d *Display) recordFlex() {
	var (
		records = createTextView("RECORDS").
			SetTextStyle(tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Bold(true))
	)

	d.RecordFlex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(records, 0, 1, false)
}

func createTextView(text string) *tview.TextView {
	return tview.NewTextView().
		SetText(text).
		SetTextColor(tcell.ColorWhite).
		SetTextAlign(tview.AlignCenter)
}
