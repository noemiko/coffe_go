package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/noemiko/coffe_go/machine"
)

type APICommandBody struct {
	Value string `json:"command"`
}

type APIResponseBody struct {
	Value string `json:"response"`
}

//Logic to get command and retrieve message from machine
func PostMethod(c *gin.Context) {
	var command APICommandBody

	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(jsonDataBytes, &command)
	if err != nil {
		fmt.Println(err)
		return
	}
	message := machine.Execute(command.Value)
	response := &APIResponseBody{
		Value: message,
	}
	json.Marshal(response)
	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()

	router.POST("/drink", PostMethod)

	listenPort := "4000"
	router.Run(":" + listenPort)
}
