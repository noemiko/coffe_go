package machine

import (
	"fmt"
	"strconv"
	"strings"
)

func parseToDrinkProperties(rawCommand string) *DrinkProps {
	splittedCommand := strings.Split(rawCommand, ":")
	var isStick bool
	lastElement := splittedCommand[2]
	if string(lastElement) == "" {
		isStick = false
	} else if string(lastElement) == "0" {
		isStick = true
	}

	amountOfSugar, err := strconv.Atoi(splittedCommand[1])
	if err != nil {
		fmt.Println(err)
	}
	drinkProps := &DrinkProps{
		Type:          splittedCommand[0],
		AmountOfSugar: amountOfSugar,
		IsStick:       isStick,
	}
	return drinkProps
}

func parseToOperationalProperties(rawCommand string) *OperationalProps {
	operationalProps := &OperationalProps{
		Type:    rawCommand[0:1],
		Details: rawCommand[2:],
	}
	return operationalProps
}
