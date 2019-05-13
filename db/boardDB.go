package db

import (
	"encoding/json"
	"fmt"

	"github.com/gueradevelopment/personal-context/models"
)

// BoardDB - Board model database accessor
type BoardDB struct{}

// Get - retrieves a single resource
func (db *BoardDB) Get(id string, c chan Result) {
	defer close(c)

	queryID := fmt.Sprintf(`{"id":"%s"}`, id)
	res := make(chan string)
	go broker.SendAndReceive(queryID, boardKey+"retrieve", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitResponse(response)
}

// GetAll - retrieves all resources
func (db *BoardDB) GetAll(c chan ResultArray, where map[string][]string) {
	defer close(c)

	userID := where["userId"][0]
	queryID := fmt.Sprintf(`{"userId":"%s"}`, userID)

	res := make(chan string)
	go broker.SendAndReceive(queryID, boardKey+"retrieveAll", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitArray(response)
}

// Add - creates a resource
func (db *BoardDB) Add(item models.Board, c chan Result) {
	defer close(c)
	marshalled, _ := json.Marshal(item) // It was unmarshalled at the controller, it should no be any error here
	res := make(chan string)
	go broker.SendAndReceive(string(marshalled), boardKey+"create", res)

	response := <-res
	fmt.Println(response)
	c <- parseRabbitResponse(response)
}

// Edit - updates a resource
func (db *BoardDB) Edit(item models.Board, c chan Result) {
	defer close(c)

	marshalled, _ := json.Marshal(item) // It was unmarshalled at the controller, it should no be any error here
	res := make(chan string)
	fmt.Println(string(marshalled))
	go broker.SendAndReceive(string(marshalled), boardKey+"update", res)

	response := <-res
	fmt.Println(response)
	c <- parseRabbitResponse(response)
}

// Delete - deletes a resource
func (db *BoardDB) Delete(id string, c chan Result) {
	defer close(c)

	queryID := fmt.Sprintf(`{"id":"%s"}`, id)
	res := make(chan string)
	go broker.SendAndReceive(queryID, boardKey+"delete", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitResponse(response)
}
