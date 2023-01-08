package users

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"snake/customerror"
	"strconv"
	"strings"
)

type Player struct {
	Name      string
	LastScore int
	BestScore int
}

const userFileName = "data/users.txt" // Имя файла с даными о игроках

// createStateFile создат файл для хранения данных о игроках
func createStateFile() {
	file, err := os.Create(userFileName)

	customerror.CheckError(customerror.CustomError{
		Message: "Что-то пошло не так при создании файла",
		Error:   err,
	})

	defer file.Close()
}

/*
GetPlayers GetUsers получает всех игроков из хранилица Возращает массив
игроков []Player
*/
func GetPlayers() []Player {
	var users []Player // массив с игроками

	file, err := os.Open(userFileName)
	customerror.CheckError(customerror.CustomError{
		Message: "Произошла ошибка с открытием файла",
		Error:   err,
	})

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// получаем игроков из файла в виде строки из трех частей
		// и разбиваем ее на тип Player
		usersInFile := strings.Split(scanner.Text(), " ")
		if len(usersInFile) != 3 { // если не совпадает с типом игрока
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

		user := Player{
			Name:      usersInFile[0],
			LastScore: lastScore,
			BestScore: bestScore,
		}
		users = append(users, user)
	}

	return users
}

func WritePlayer(player Player) {
	playerToString := fmt.Sprintf(
		"%v %v %v\n", player.Name, player.LastScore, player.BestScore,
	)

	file, err := os.OpenFile(userFileName, os.O_APPEND|os.O_WRONLY, 0600)
	customerror.CheckError(customerror.CustomError{
		Message: "Что-то пошло не так при открытии файла",
		Error:   err,
	})
	defer file.Close()

	if _, err = file.WriteString(playerToString); err != nil {
		panic(customerror.ErrorString(customerror.CustomError{
			Message: "Что-то пошло не так во время записи в файл",
			Error:   err,
		}))
	}
}

// FindPlayer ищет игрока по имени или возращает ошибку
// Возращает второй аргумент bool
func FindPlayer(name string) (Player, bool) {
	players := GetPlayers()

	for _, p := range players {
		if p.Name == name {
			return p, true
		}
	}

	return Player{}, false
}

/*
NewPlayer Интерфейс для структуры Player, принимает имя (name string) нового
пользователя. Возращает новго игрока типа Player
*/
func NewPlayer(name string) Player {
	if _, err := os.Stat(userFileName); errors.Is(err, os.ErrNotExist) {
		createStateFile()
	}

	if _, p := FindPlayer(name); p {
		panic(customerror.ErrorString(customerror.CustomError{
			Message: "Пользователь с таким именем уже сущетвует",
		}))
	}

	return Player{Name: name, LastScore: 0, BestScore: 0}
}
