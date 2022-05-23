package machine

import (
	"fmt"
)

func getDrinkMessage(drinkProp *DrinkProps) string {
	message := getDrinkTypeMessage(drinkProp)
	if message == UnknownCommandMessage {
		return UnknownCommandMessage
	}

	message += getDrinkAmounthOfSugarMessage(drinkProp)
	message += getIsStickMessage(drinkProp)

	return message
}

func getOperationalMessage(operationalProp *OperationalProps) string {
	if operationalProp.Type == "M" && operationalProp.Details == "message-content" {
		return "Drink maker forwards any message received onto the coffee machine interface for the customer to see"
	}
	return UnknownCommandMessage
}

func getDrinkTypeMessage(drinkProp *DrinkProps) string {
	Tea := "T"
	Coffe := "C"
	Chocolate := "H"

	if drinkProp.Type == Tea {
		return "Drink maker makes 1 tea "
	} else if drinkProp.Type == Chocolate {
		return "Drink maker makes 1 chocolate "
	} else if drinkProp.Type == Coffe {
		return "Drink maker makes 1 coffee "
	} else {
		return UnknownCommandMessage
	}
}

func getDrinkAmounthOfSugarMessage(drinkProp *DrinkProps) string {
	if drinkProp.AmountOfSugar == 0 {
		return "with no sugar "
	} else if drinkProp.AmountOfSugar == 1 {
		return fmt.Sprintf("with %d sugar ", drinkProp.AmountOfSugar)
	} else {
		return fmt.Sprintf("with %d sugars ", drinkProp.AmountOfSugar)
	}
}

func getIsStickMessage(drinkProp *DrinkProps) string {
	if drinkProp.IsStick {
		return "and a stick"
	} else {
		if drinkProp.AmountOfSugar == 0 {
			return "- and therefore no stick"

		} else {
			return "- and no stick"
		}
	}
}
