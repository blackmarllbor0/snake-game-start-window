package main

import (
	"snake/display"
	"snake/players"
)

func main() {
	players.CheckDir()
	display := display.Display{}
	display.Displaying()
}
