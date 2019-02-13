package db

import (
	"errors"

	"github.com/gueradevelopment/personal-context/models"
)

// BoardDB - Board model database accessor
type BoardDB struct{}

var (
	boardItems = make(map[string]models.Board)
)

// Get - retrieves a single resource
func (db *BoardDB) Get(id string, c chan Result) {
	defer close(c)
	result := Result{}
	for ID, item := range boardItems {
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
func (db *BoardDB) GetAll(c chan ResultArray, where map[string][]string) {
	defer close(c)
	result := ResultArray{}
	var arr = []Model{}
	var guerabookID string
	if where["guerabookId"] != nil {
		guerabookID = where["guerabookId"][0]
	}
	for _, v := range boardItems {
		if guerabookID != "" && v.GuerabookID == guerabookID {
			arr = append(arr, v)
		}
		if guerabookID == "" {
			arr = append(arr, v)
		}
	}
	result.Result = arr
	c <- result
}

// Add - creates a resource
func (db *BoardDB) Add(item models.Board, c chan Result) {
	defer close(c)
	result := Result{}
	if boardItems[item.ID] == (models.Board{}) {
		boardItems[item.ID] = item
		result.Result = item
	} else {
		result.Err = errors.New("Duplicated ID")
	}
	c <- result
}

// Edit - updates a resource
func (db *BoardDB) Edit(item models.Board, c chan Result) {
	defer close(c)
	result := Result{}
	if boardItems[item.ID] == (models.Board{}) {
		result.Err = errors.New("No such ID")
	} else {
		boardItems[item.ID] = item
		result.Result = item
	}
	c <- result
}

// Delete - deletes a resource
func (db *BoardDB) Delete(id string, c chan Result) {
	defer close(c)
	result := Result{}
	item := boardItems[id]
	if item == (models.Board{}) {
		result.Err = errors.New("No such ID")
	} else {
		result.Result = item
		delete(boardItems, id)
	}
	c <- result
}
