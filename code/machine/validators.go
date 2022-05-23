package machine

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type DrinkCommand struct {
	Drink string `validate:"required,oneof=C H T"`
	Sugar string `validate:"omitempty,numeric"`
	Stick string `validate:"oneof=0 ''"`
}

func validateDrinkCommand(rawCommand string) string {
	splittedCommand := strings.Split(rawCommand, ":")
	drinkCommand := &DrinkCommand{
		Drink: splittedCommand[0],
		Sugar: splittedCommand[1],
		Stick: splittedCommand[2],
	}
	validate := validator.New()
	err := validate.Struct(drinkCommand)
	if err == nil {
		return ""
	}
	message := ""
	for _, err := range err.(validator.ValidationErrors) {
		message += strings.ToLower(err.Field()) + ", "
	}
	return strings.Trim(message, ", ")
}
