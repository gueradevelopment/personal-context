package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gueradevelopment/personal-context/services"
)

var (
	broker       = services.RabbitServiceInit()
	guerabookKey = "personal.guerabook."
	boardKey     = "personal.board."
	checklistKey = "personal.checklist."
	taskKey      = "personal.task."
)

// Model interface to wrap data types
type Model interface{}

// Result struct to wrap channel tuple
type Result struct {
	Result Model
	Err    error
}

// ResultArray struct to wrap channel tuple
type ResultArray struct {
	Result []Model
	Err    error
}

// Database interface to wrap data accessors
type Database interface {
	Get(string, chan Result)
	GetAll(chan ResultArray, map[string][]string)
	Add(Model, chan Result)
	Edit(Model, chan Result)
	Delete(string, chan Result)
}

func parseRabbitResponse(response string) Result {
	result := Result{}
	responseMap := make(map[string]interface{})
	json.Unmarshal([]byte(response), &responseMap)

	if responseMap["type"] == "success" {
		result.Result = responseMap["data"]
	} else {
		result.Err = errors.New(responseMap["reason"].(string))
	}

	fmt.Println(result.Err)

	return result
}

func parseRabbitArray(response string) ResultArray {
	result := ResultArray{}

	responseMap := make(map[string]interface{})
	json.Unmarshal([]byte(response), &responseMap)

	if responseMap["type"] == "success" {
		result.Result = []Model{responseMap["data"]}
	} else {
		result.Err = errors.New("Server error")
	}

	return result
}
