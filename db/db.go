package db

import (
	"encoding/json"
	"errors"
	"net/http"

	"personal-context/services"
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
	Code   int
}

// ResultArray struct to wrap channel tuple
type ResultArray struct {
	Result []Model
	Err    error
	Code   int
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
		result.Code = http.StatusOK
	} else {
		result.Err = errors.New(responseMap["reason"].(string))
		var code int

		switch responseMap["errorType"] {
		case "NotFoundException":
			code = http.StatusNotFound
			break
		case "UnsupportedActionException":
			code = http.StatusNotImplemented
			break
		case "BadRequestException":
			code = http.StatusBadRequest
			break
		default:
			code = http.StatusInternalServerError
			break
		}

		result.Code = code
	}

	return result
}

func parseRabbitArray(response string) ResultArray {
	result := ResultArray{}

	responseMap := make(map[string]interface{})
	json.Unmarshal([]byte(response), &responseMap)

	if responseMap["type"] == "success" {
		result.Result = []Model{responseMap["data"]}
		result.Code = http.StatusOK
	} else {
		result.Err = errors.New(responseMap["reason"].(string))
		result.Code = http.StatusNotFound
	}

	return result
}
