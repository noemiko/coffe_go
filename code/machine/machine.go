package machine

import (
	"strings"
)

type DrinkProps struct {
	Type          string
	AmountOfSugar int
	IsStick       bool
}
type OperationalProps struct {
	Type    string
	Details string
}


const (
	DrinkCommandType       string = "DrinkCommand"
	OperationalCommandType string = "OperationalCommand"
	UnknownCommandType     string = "UnknowCommand"

	UnknownCommandMessage string = "Unknown command"
)

//Recognize type of command based on numer of :
func getTypeOfCommand(rawCommand string) string {
	splittedCommand := strings.Split(rawCommand, ":")
	if len(splittedCommand) == 3 {
		return DrinkCommandType
	} else if len(splittedCommand) == 2 {
		return OperationalCommandType
	}
	return UnknownCommandType
}


//Recognize type of command and send proper message
func Execute(rawCommand string) string {
	commandType := getTypeOfCommand(rawCommand)
	if commandType == DrinkCommandType {
		invalidFields := validateDrinkCommand(rawCommand)
		if invalidFields != "" {
			return "Unknown information about " + invalidFields
		}

		drinkProp := parseResponseToDrinkProperties(rawCommand)
		return getDrinkMessage(drinkProp)

	} else if commandType == OperationalCommandType {
		operationalProp := parseResponseToOperationalProperties(rawCommand)
		return getOperationalMessage(operationalProp)
	} else {
		return UnknownCommandMessage
	}
}
