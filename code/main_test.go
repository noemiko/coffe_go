package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func sendToDrinkMachineCommand(command string) (int, string) {
	requestBody := &APICommandBody{
		Value: command,
	}

	data, _ := json.Marshal(requestBody)
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/drink", PostMethod)

	req, err := http.NewRequest(http.MethodPost, "/drink", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Couldn't create request: %v \n", err)
	}
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var responseBody APIResponseBody
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	return response.Code, responseBody.Value
}

func TestDrinksMachineCommands(t *testing.T) {

	var machineUseCases = []struct {
		in          string
		expectedOut string
	}{
		{"T:1:0", "Drink maker makes 1 tea with 1 sugar and a stick"},
		{"T:2:0", "Drink maker makes 1 tea with 2 sugars and a stick"},
		{"T:2:", "Drink maker makes 1 tea with 2 sugars - and no stick"},
		{"H::", "Drink maker makes 1 chocolate with no sugar - and therefore no stick"},
		{"H:0:", "Drink maker makes 1 chocolate with no sugar - and therefore no stick"},
		{"H:10:", "Drink maker makes 1 chocolate with 10 sugars - and no stick"},
		{"C:2:0", "Drink maker makes 1 coffee with 2 sugars and a stick"},

		{"H:10:foo", "Unknown information about stick"},
		{"o:jj: jj", "Unknown information about drink, sugar, stick"},
		{"H:ggg:", "Unknown information about sugar"},
		{"H:ggg:ggg", "Unknown information about sugar, stick"},
		{"L::", "Unknown information about drink"},
		{"H:2:1:5", "Unknown command"},

		{"C", "Unknown command"},

		{":::foo", "Unknown command"},
		{"", "Unknown command"},
		{"foo::::", "Unknown command"},
		{"give me coffe", "Unknown command"},
	}

	for index, useCase := range machineUseCases {
		fmt.Printf("Case %d with command %s \n", index+1, useCase.in)
		responseCode, messageResponse := sendToDrinkMachineCommand(useCase.in)
		assert.Equal(t, 200, responseCode)
		assert.Equal(t, useCase.expectedOut, messageResponse)
	}
}
