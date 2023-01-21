package players

import (
	"strings"
)

type Player struct {
	name      string
	lastScore int
	bestScore int
}

// NewPlayer Интерфейс для структуры Player
func NewPlayer(name string) *Player {
	// вырезаем все пробелы из имени игрока
	userName := strings.Replace(name, " ", "_", -1)
	return &Player{name: userName, lastScore: 0, bestScore: 0}
}

// GetName получить имя
func (p *Player) GetName() string {
	return p.name
}

// SetName изменить имя
func (p *Player) SetName(name string) {
	p.name = name
}

// GetLastScore получить последий счет
func (p *Player) GetLastScore() int {
	return p.lastScore
}

// SetLastScore изменить последний счет
func (p *Player) SetLastScore(lastScore int) {
	p.lastScore = lastScore
}

// GetBestScore получить лучший счет
func (p *Player) GetBestScore() int {
	return p.bestScore
}

// SetBestScore изменить лучший счет
func (p *Player) SetBestScore(bestScore int) {
	p.bestScore = bestScore
}
