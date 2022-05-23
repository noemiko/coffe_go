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
	UnknownCommandMessage string = "Unknown command"
)

type Command struct {
	raw        string
	structured interface{} // OperationalProps or DrinkProps
	message    string
}

func newDrinkCommand(rawCommand string) *Command {
	invalidFields := validateDrinkCommand(rawCommand)
	message := ""

	drinkProp := parseToDrinkProperties(rawCommand)

	if invalidFields != "" {
		message = "Unknown information about " + invalidFields
	} else {
		message = getDrinkMessage(drinkProp)
	}
	return &Command{
		raw:        rawCommand,
		structured: drinkProp,
		message:    message,
	}

}

func newOperationalCommand(rawCommand string) *Command {
	operationalProp := parseToOperationalProperties(rawCommand)
	message := getOperationalMessage(operationalProp)
	return &Command{
		raw:        rawCommand,
		structured: operationalProp,
		message:    message,
	}
}

// in Unknown Comand is hard to recognize structure
// that why it is empty
func newUnknownCommand(rawCommand string) *Command {
	return &Command{
		raw:        rawCommand,
		structured: "",
		message:    UnknownCommandMessage,
	}
}

func (c Command) getMessage() string {
	return c.message
}

// Fabric with commands
// Recognize type of command based on numer of :
func getProperCommandStruct(rawCommand string) *Command {
	splittedCommand := strings.Split(rawCommand, ":")
	if len(splittedCommand) == 3 {
		return newDrinkCommand(rawCommand)
	} else if len(splittedCommand) == 2 {
		return newOperationalCommand(rawCommand)
	}
	return newUnknownCommand(rawCommand)
}

// run machine and get response about result
func Execute(rawCommand string) string {
	updateHistory(rawCommand)
	commandStruct := getProperCommandStruct(rawCommand)
	return commandStruct.getMessage()
}
