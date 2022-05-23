package machine

import (
	"strings"
)

var history string

func updateHistory(command string) string {
	history += command + ", "
	return history
}

func GetComandsHistory() string {
	return strings.Trim(history, ", ")
}
