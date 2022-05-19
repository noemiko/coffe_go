package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "fmt"
  "bytes"
  "github.com/gin-gonic/gin"
  "github.com/stretchr/testify/assert"
  "encoding/json"


)

// func TestPostMethod(t *testing.T) {
//     // Switch to test mode so you don't get such noisy output
//     gin.SetMode(gin.TestMode)

//     // Setup your router, just like you did in your main function, and
//     // register your routes
//     r := gin.Default()
//     r.POST("/", PostMethod)

//     // Create the mock request you'd like to test. Make sure the second argument
//     // here is the same as one of the routes you defined in the router setup
//     // block!
//     req, err := http.NewRequest(http.MethodPost, "/", nil)
//     if err != nil {
//         t.Fatalf("Couldn't create request: %v\n", err)
//     }

//     // Create a response recorder so you can inspect the response
//     w := httptest.NewRecorder()

//     // Perform the request
//     r.ServeHTTP(w, req)
//     fmt.Println(w.Body)

//     // Check to see if the response was what you expected
//     if w.Code == http.StatusOK {
//         t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
//     } else {
//         t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
//     }
// }

func TestGetMethod(t *testing.T) {
    // Switch to test mode so you don't get such noisy output
    gin.SetMode(gin.TestMode)

    // Setup your router, just like you did in your main function, and
    // register your routes
    r := gin.Default()
    r.GET("/", GetMethod)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
    req, err := http.NewRequest(http.MethodGet, "/", nil)
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    // Create a response recorder so you can inspect the response
    w := httptest.NewRecorder()

    // Perform the request
    r.ServeHTTP(w, req)
    fmt.Println(w.Body)

    // Check to see if the response was what you expected
    if w.Code == http.StatusOK {
        t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
    } else {
        t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
    }
}
// func getRegistrationPOSTPayload() string {
// 	params := url.Values{}
// 	params.Add("username", "u1")
// 	params.Add("password", "p1")

// 	return params.Encode()
// }

func TestTeaMaker(t *testing.T) {
        command := &CommandBody{
            Value : "T:1:0",
	    }
        expectBody:= "\"Drink maker makes 1 tea with 1 sugar and a stick\""
        data, _ := json.Marshal(command)
        gin.SetMode(gin.TestMode)
        r := gin.Default()
        r.POST("/drink", PostMethod)
    
        // Create the mock request you'd like to test. Make sure the second argument
        // here is the same as one of the routes you defined in the router setup
        // block!
        req, err := http.NewRequest(http.MethodPost, "/drink", bytes.NewBuffer(data))
        if err != nil {
            t.Fatalf("Couldn't create request: %v\n", err)
        }
    
        // Create a response recorder so you can inspect the response
        w := httptest.NewRecorder()
    
        // Perform the request
        r.ServeHTTP(w, req)
        fmt.Println(w.Body)

	    assert.Equal(t, 200, w.Code)
	    assert.Equal(t, expectBody, w.Body.String())
  
}

func TestCoffeMaker(t *testing.T) {
    command := &CommandBody{
        Value : "C:1:0",
    }
    expectBody:= "\"Drink maker makes 1 coffee with 2 sugars and a stick\""
    data, _ := json.Marshal(command)
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.POST("/drink", PostMethod)

    // Create the mock request you'd like to test. Make sure the second argument
    // here is the same as one of the routes you defined in the router setup
    // block!
    req, err := http.NewRequest(http.MethodPost, "/drink", bytes.NewBuffer(data))
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    fmt.Println(w.Body)

    assert.Equal(t, 200, w.Code)
    assert.Equal(t, expectBody, w.Body.String())

}

func TestCHocolateMaker(t *testing.T) {
    command := &CommandBody{
        Value : "H:1:0",
    }
    expectBody:= "\"Drink maker makes 1 chocolate with no sugar - and therefore no stick\""
    data, _ := json.Marshal(command)
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.POST("/drink", PostMethod)

    req, err := http.NewRequest(http.MethodPost, "/drink", bytes.NewBuffer(data))
    if err != nil {
        t.Fatalf("Couldn't create request: %v\n", err)
    }

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    fmt.Println(w.Body)

    assert.Equal(t, 200, w.Code)
    assert.Equal(t, expectBody, w.Body.String())

}



