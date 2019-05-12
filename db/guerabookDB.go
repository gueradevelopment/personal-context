package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gueradevelopment/personal-context/models"
	"github.com/gueradevelopment/personal-context/services"
)

// GuerabookDB - Guerabook model database accessor
type GuerabookDB struct{}

var (
	guerabookItems = make(map[string]models.Guerabook)
	broker         = services.RabbitServiceInit()
	routeKey       = "personal.guerabook."
)

// Get - retrieves a single resource
func (db *GuerabookDB) Get(id string, c chan Result) {
	defer close(c)

	queryID := fmt.Sprintf(`{"id":"%s"}`, id)
	res := make(chan string)
	go broker.SendAndReceive(queryID, routeKey+"retrieve", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitResponse(response)
}

// GetAll - retrieves all resources
func (db *GuerabookDB) GetAll(c chan ResultArray, where map[string][]string) {
	defer close(c)

	userID := where["userId"][0]
	queryID := fmt.Sprintf(`{"userId":"%s"}`, userID)

	res := make(chan string)
	go broker.SendAndReceive(queryID, routeKey+"retrieveAll", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitArray(response)
}

// Add - creates a resource
func (db *GuerabookDB) Add(item models.Guerabook, c chan Result) {
	defer close(c)
	marshalled, _ := json.Marshal(item) // It was unmarshalled at the controller, it should no be any error here
	res := make(chan string)
	go broker.SendAndReceive(string(marshalled), routeKey+"create", res)

	response := <-res
	fmt.Println(response)
	c <- parseRabbitResponse(response)
}

// Edit - updates a resource
func (db *GuerabookDB) Edit(item models.Guerabook, c chan Result) {
	defer close(c)

	marshalled, _ := json.Marshal(item) // It was unmarshalled at the controller, it should no be any error here
	res := make(chan string)
	go broker.SendAndReceive(string(marshalled), routeKey+"update", res)

	response := <-res
	fmt.Println(response)
	c <- parseRabbitResponse(response)
}

// Delete - deletes a resource
func (db *GuerabookDB) Delete(id string, c chan Result) {
	defer close(c)

	queryID := fmt.Sprintf(`{"id":"%s"}`, id)
	res := make(chan string)
	go broker.SendAndReceive(queryID, routeKey+"delete", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitResponse(response)
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
