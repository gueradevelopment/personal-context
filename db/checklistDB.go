package db

import (
	"encoding/json"
	"fmt"

	"github.com/gueradevelopment/personal-context/models"
)

// ChecklistDB - Checklist model database accessor
type ChecklistDB struct{}

// Get - retrieves a single resource
func (db *ChecklistDB) Get(id string, c chan Result) {
	defer close(c)

	queryID := fmt.Sprintf(`{"id":"%s"}`, id)
	res := make(chan string)
	go broker.SendAndReceive(queryID, checklistKey+"retrieve", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitResponse(response)
}

// GetAll - retrieves all resources
func (db *ChecklistDB) GetAll(c chan ResultArray, where map[string][]string) {
	defer close(c)

	userID := where["userId"][0]
	queryID := fmt.Sprintf(`{"userId":"%s"}`, userID)

	res := make(chan string)
	go broker.SendAndReceive(queryID, checklistKey+"retrieveAll", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitArray(response)
}

// Add - creates a resource
func (db *ChecklistDB) Add(item models.Checklist, c chan Result) {
	defer close(c)
	marshalled, _ := json.Marshal(item) // It was unmarshalled at the controller, it should no be any error here
	res := make(chan string)
	go broker.SendAndReceive(string(marshalled), checklistKey+"create", res)

	response := <-res
	fmt.Println(response)
	c <- parseRabbitResponse(response)
}

// Edit - updates a resource
func (db *ChecklistDB) Edit(item models.Checklist, c chan Result) {
	defer close(c)

	marshalled, _ := json.Marshal(item) // It was unmarshalled at the controller, it should no be any error here
	res := make(chan string)
	go broker.SendAndReceive(string(marshalled), checklistKey+"update", res)

	response := <-res
	fmt.Println(response)
	c <- parseRabbitResponse(response)
}

// Delete - deletes a resource
func (db *ChecklistDB) Delete(id string, c chan Result) {
	defer close(c)

	queryID := fmt.Sprintf(`{"id":"%s"}`, id)
	res := make(chan string)
	go broker.SendAndReceive(queryID, checklistKey+"delete", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitResponse(response)
}
