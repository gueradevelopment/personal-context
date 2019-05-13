package db

import (
	"encoding/json"
	"fmt"

	"github.com/gueradevelopment/personal-context/models"
)

// TaskDB - Task model database accessor
type TaskDB struct{}

// Get - retrieves a single resource
func (db *TaskDB) Get(id string, c chan Result) {
	defer close(c)

	queryID := fmt.Sprintf(`{"id":"%s"}`, id)
	res := make(chan string)
	go broker.SendAndReceive(queryID, taskKey+"retrieve", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitResponse(response)
}

// GetAll - retrieves all resources
func (db *TaskDB) GetAll(c chan ResultArray, where map[string][]string) {
	defer close(c)

	userID := where["userId"][0]
	queryID := fmt.Sprintf(`{"userId":"%s"}`, userID)

	res := make(chan string)
	go broker.SendAndReceive(queryID, taskKey+"retrieveAll", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitArray(response)
}

// Add - creates a resource
func (db *TaskDB) Add(item models.Task, c chan Result) {
	defer close(c)
	marshalled, _ := json.Marshal(item) // It was unmarshalled at the controller, it should no be any error here
	res := make(chan string)
	go broker.SendAndReceive(string(marshalled), taskKey+"create", res)

	response := <-res
	fmt.Println(response)
	c <- parseRabbitResponse(response)
}

// Edit - updates a resource
func (db *TaskDB) Edit(item models.Task, c chan Result) {
	defer close(c)

	marshalled, _ := json.Marshal(item) // It was unmarshalled at the controller, it should no be any error here
	res := make(chan string)
	go broker.SendAndReceive(string(marshalled), taskKey+"update", res)

	response := <-res
	fmt.Println(response)
	c <- parseRabbitResponse(response)
}

// Delete - deletes a resource
func (db *TaskDB) Delete(id string, c chan Result) {
	defer close(c)

	queryID := fmt.Sprintf(`{"id":"%s"}`, id)
	res := make(chan string)
	go broker.SendAndReceive(queryID, taskKey+"delete", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitResponse(response)
}
