package main

import (
	"fmt"
	"os"
	"snake/display"
	"snake/players"
)

func main() {
	players.CheckDir()
	display := display.Display{}

	// будет выбора запуска игры через два вида интерфейса
	if len(os.Args) <= 1 {
		// для запуска игры через граф. интерфейс
		display.Displaying()
	} else {
		// для запуск игры через передачу аргументов
		fmt.Println(os.Args[1:])
	}
}
