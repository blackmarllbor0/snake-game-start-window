package display

import (
	"snake/customerror"
	"snake/players"

	"github.com/gdamore/tcell/v2"
	tview "github.com/rivo/tview"
)

type Display struct {
	app       *tview.Application
	mainBlock *tview.Flex
	wall      *tview.Box
	GameDisplay
	RecordDisplay
}

// константы для хранения цвета
const (
	cw = tcell.ColorWhite
	cb = tcell.ColorBlack
)

// Displaying функция запуска графического приложения с главного меню
func (d *Display) Displaying() {
	d.app = tview.NewApplication() // Графическое приложение

	d.remote() // пульт управления нажатий клавиш

	// запуск приложения
	if err := d.app.SetRoot(d.initMain(), true).Run(); err != nil {
		panic(customerror.ErrorString(customerror.CustomError{
			Message: "Приложение упало при запуске",
			Error:   err,
		}))
	}
}

// initMain инициализация главного блока
func (d *Display) initMain() *tview.Flex {
	d.mainBlock = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(d.GameDisplay.init(), 0, 2, true).
		AddItem(d.initWall(), 1, 0, false).
		AddItem(d.RecordDisplay.init(), 0, 1, false)
	return d.mainBlock
}

// initWall стенка для разделения основных блоков
func (d Display) initWall() *tview.Box {
	d.wall = tview.NewBox().SetBackgroundColor(cw)
	return d.wall
}

// remote пульт управления нажатий
func (d *Display) remote() {
	d.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' && !d.GameDisplay.inputName.HasFocus() {
			d.app.Stop()
		} else if event.Key() == tcell.KeyEnter && d.GameDisplay.inputName.HasFocus() {
			d.app.SetFocus(d.RecordDisplay.block)
		} else if event.Key() == tcell.KeyTAB {
			d.app.SetFocus(d.GameDisplay.inputName)
		} else if event.Rune() == 's' {
			// при запуске игры сохроняем имя игрока в файл
			players.WritePlayer(players.NewPlayer(d.GameDisplay.inputName.GetText()))
		}
		return event
	})
}
