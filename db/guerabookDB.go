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
	res := make(chan string)
	go broker.SendAndReceive("{\"title\": \"I am the Senate\"}", "guerabook.create", res)
	response := <-res
	fmt.Println(response)

	result := Result{}
	for ID, item := range guerabookItems {
		if ID == id {
			result.Result = item
			result.Err = nil
			break
		}
	}
	if result.Result == nil {
		result.Err = errors.New("No result")
	}
	c <- result
}

// GetAll - retrieves all resources
func (db *GuerabookDB) GetAll(c chan ResultArray, where map[string][]string) {
	defer close(c)
	result := ResultArray{}
	var arr = []Model{}
	for _, v := range guerabookItems {
		arr = append(arr, v)
	}
	result.Result = arr
	c <- result
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
	result := Result{}
	if guerabookItems[item.ID] == (models.Guerabook{}) {
		result.Err = errors.New("No such ID")
	} else {
		guerabookItems[item.ID] = item
		result.Result = item
	}
	c <- result
}

// Delete - deletes a resource
func (db *GuerabookDB) Delete(id string, c chan Result) {
	defer close(c)
	result := Result{}
	item := guerabookItems[id]
	if item == (models.Guerabook{}) {
		result.Err = errors.New("No such ID")
	} else {
		result.Result = item
		delete(guerabookItems, id)
	}
	c <- result
}

func parseRabbitResponse(response string) Result {
	result := Result{}
	responseMap := make(map[string]interface{})
	json.Unmarshal([]byte(response), &responseMap)

	if responseMap["type"] == "success" {
		result.Result = responseMap["data"]
	} else {
		result.Err = errors.New("Server error")
	}

	return result
}
