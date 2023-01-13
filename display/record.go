package display

import (
	"fmt"
	"snake/players"
	"sort"

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
	var flex = tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(r.createList(), 0, 4, false).
		AddItem(tview.NewBox(), 0, 1, false)

	r.block = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(r.createLabel(), 0, 1, false).
		AddItem(flex, 0, 3, false)
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

func (r *RecordDisplay) createList() *tview.List {
	r.bestPlayers = tview.NewList()
	r.bestPlayers.SetTitleAlign(tview.AlignCenter)

	var playerList = players.GetPlayers()

	sort.Slice(playerList, func(i, j int) bool {
		return playerList[i].GetBestScore() > playerList[j].GetBestScore()
	})

	if len(playerList) >= 5 {
		for i := 0; i < 6; i++ {
			var currentPlayer = playerList[i]
			if currentPlayer.GetBestScore() > 0 {
				var player = fmt.Sprintf(
					"%s - %d\n", currentPlayer.GetName(), currentPlayer.GetBestScore(),
				)
				r.bestPlayers.AddItem(player, "", tcell.RuneVLine, nil)
			}
		}
	} else {
		for i := 0; i < len(playerList); i++ {
			var currentPlayer = playerList[i]
			if currentPlayer.GetBestScore() > 0 {
				var player = fmt.Sprintf(
					"%s - %d\n", currentPlayer.GetName(), currentPlayer.GetBestScore(),
				)
				r.bestPlayers.AddItem(player, "", tcell.RuneVLine, nil)
			}
		}
	}

	return r.bestPlayers
}
