package db

import (
	"encoding/json"
	"fmt"

	"github.com/gueradevelopment/personal-context/models"
)

// GuerabookDB - Guerabook model database accessor
type GuerabookDB struct{}

// Get - retrieves a single resource
func (db *GuerabookDB) Get(id string, c chan Result) {
	defer close(c)

	queryID := fmt.Sprintf(`{"id":"%s"}`, id)
	res := make(chan string)
	go broker.SendAndReceive(queryID, guerabookKey+"retrieve", res)

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
	go broker.SendAndReceive(queryID, guerabookKey+"retrieveAll", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitArray(response)
}

// Add - creates a resource
func (db *GuerabookDB) Add(item models.Guerabook, c chan Result) {
	defer close(c)
	marshalled, _ := json.Marshal(item) // It was unmarshalled at the controller, it should no be any error here
	res := make(chan string)
	go broker.SendAndReceive(string(marshalled), guerabookKey+"create", res)

	response := <-res
	fmt.Println(response)
	c <- parseRabbitResponse(response)
}

// Edit - updates a resource
func (db *GuerabookDB) Edit(item models.Guerabook, c chan Result) {
	defer close(c)

	marshalled, _ := json.Marshal(item) // It was unmarshalled at the controller, it should no be any error here
	res := make(chan string)
	go broker.SendAndReceive(string(marshalled), guerabookKey+"update", res)

	response := <-res
	fmt.Println(response)
	c <- parseRabbitResponse(response)
}

// Delete - deletes a resource
func (db *GuerabookDB) Delete(id string, c chan Result) {
	defer close(c)

	queryID := fmt.Sprintf(`{"id":"%s"}`, id)
	res := make(chan string)
	go broker.SendAndReceive(queryID, guerabookKey+"delete", res)

	response := <-res
	fmt.Println(response)

	c <- parseRabbitResponse(response)
}
