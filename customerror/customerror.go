package customerror

import (
	"fmt"
)

type CustomError struct {
	Message string
	Error   error
}

// CheckError проверяет не равна ли ошибка nil, а если равно, то создает панику
func CheckError(customError CustomError) {
	if customError.Error != nil {
		panic(ErrorString(customError))
	}
}

// ErrorString возращает отформатировану строку с кастомной ошибкой
func ErrorString(customError CustomError) string {
	return fmt.Sprintf("%v\n%v\n", customError.Message, customError.Error)
}
