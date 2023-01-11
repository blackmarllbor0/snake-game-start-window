package display

import "github.com/rivo/tview"

type Game struct {
	block *tview.Flex
}

func (g *Game) init() *tview.Flex {
	g.block = tview.NewFlex()

	return g.block
}
