package players

import (
	"bufio"
	"fmt"
	"os"
	"snake/customerror"
	"strconv"
	"strings"
)

// userFileName имя файла с даными о игроках
const userFileName = "data/players.txt"

// createStateFile создат файл для хранения данных о игроках
func createStateFile() {
	file, err := os.Create(userFileName)

	customerror.CheckError(customerror.CustomError{
		Message: "Что-то пошло не так при создании файла",
		Error:   err,
	})

	defer file.Close()
}

// GetPlayers GetUsers получает всех игроков из хранилица Возращает массив
// игроков []Player
func GetPlayers() []*Player {
	var (
		players   []*Player // массив с игроками
		file, err = os.Open(userFileName)
		scanner   = bufio.NewScanner(file)
	)

	customerror.CheckError(customerror.CustomError{
		Message: "Произошла ошибка с открытием файла",
		Error:   err,
	})

	customerror.CheckError(customerror.CustomError{
		Message: "Во время чтения файла произошла ошибка",
		Error:   scanner.Err(),
	})

	defer file.Close()

	for scanner.Scan() {
		// получаем игроков из файла в виде строки из трех частей
		// и разбиваем ее на тип Player
		var usersInFile = strings.Split(scanner.Text(), " ")

		// если не совпадает с типом игрока
		if len(usersInFile) != 3 {
			panic(customerror.ErrorString(customerror.CustomError{
				Message: "Пользователь состоит из трех частей",
			}))
		}

		// приобруупреобразуем строку в число
		lastScore, err := strconv.Atoi(usersInFile[1])
		customerror.CheckError(customerror.CustomError{
			Message: "Счет должен быть числовым типом данныx",
			Error:   err,
		})

		// приобруупреобразуем строку в число
		bestScore, err := strconv.Atoi(usersInFile[2])
		customerror.CheckError(customerror.CustomError{
			Message: "Рекорд должен быть числовым типом данныx",
			Error:   err,
		})

		// создаем нового игрока
		var user = &Player{
			name:      usersInFile[0],
			lastScore: lastScore,
			bestScore: bestScore,
		}
		players = append(players, user)
	}

	return players
}

// WritePlayer Записывает игрока в файл
func WritePlayer(player *Player) {
	var (
		// преобразуем все данные в одну строку
		playerToString = fmt.Sprintf(
			"%v %v %v\n", player.name, player.lastScore, player.bestScore,
		)
		// открываем файл для записи
		file, err = os.OpenFile(userFileName, os.O_APPEND|os.O_WRONLY, 0600)
	)

	customerror.CheckError(customerror.CustomError{
		Message: "Что-то пошло не так при открытии файла",
		Error:   err,
	})

	defer file.Close()

	if _, err := file.WriteString(playerToString); err != nil {
		panic(customerror.ErrorString(customerror.CustomError{
			Message: "Что-то пошло не так во время записи в файл",
			Error:   err,
		}))
	}
}

// FindPlayer ищет игрока по имени и озращает второй аргумент типа bool
func FindPlayer(name string) (*Player, bool) {
	var players = GetPlayers()

	for i := 0; i < len(players); i++ {
		if players[i].name == name {
			return players[i], true
		}
	}

	return &Player{}, false
}
