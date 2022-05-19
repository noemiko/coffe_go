package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"

)

type CommandBody struct {
    Value string `json:"command"`
}


func PostMethod(c *gin.Context) {
	// var msg CommandBody
	var command CommandBody
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)
    json.Unmarshal(jsonDataBytes, &command)
	   if err != nil {
		   return
		//    log.Fatal(err)
	   }
	
	fmt.Println("\napi.go 'PostMethod' called"+ command.Value)
	// teaCommand := "T:1:0"
	// chocolateCommand := "H::"
	// coffeCommand := "C:2:0"
	// messageContent := "M:message-content"
	orderType:= command.Value[0:1]
	message:= ""
	if orderType == "T" {
		message = "Drink maker makes 1 tea with 1 sugar and a stick"
	}else if orderType == "H" {
		message = "Drink maker makes 1 chocolate with no sugar - and therefore no stick"
	}else if orderType == "C" {
		message = "Drink maker makes 1 coffee with 2 sugars and a stick"
	}else if orderType == "M" {
		message = "Drink maker forwards any message received onto the coffee machine interface for the customer to see"
	}else{
		message = "Unknown command"
	}
	c.JSON(http.StatusOK, message)
}

func GetMethod(c *gin.Context) {
	fmt.Println("\napi.go 'GetMethod' called")
	message := "GetMethod called"
	c.JSON(http.StatusOK, message)
}

func main() {
	router := gin.Default()

	router.POST("/drink", PostMethod)
	router.GET("/", GetMethod)

	listenPort := "4000"
	// Listen and Server on the LocalHost:Port
	router.Run(":" + listenPort)
}
