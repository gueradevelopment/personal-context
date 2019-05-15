package db

import (
	"errors"

	"github.com/gueradevelopment/personal-context/models"
)

// GuerabookDB - Guerabook model database accessor
type GuerabookDB struct{}

var (
	guerabookItems = make(map[string]models.Guerabook)
)

// Get - retrieves a single resource
func (db *GuerabookDB) Get(id string, c chan Result) {
	defer close(c)
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
	result := Result{}
	if guerabookItems[item.ID] == (models.Guerabook{}) {
		guerabookItems[item.ID] = item
		result.Result = item
	} else {
		result.Err = errors.New("Duplicated ID")
	}
	c <- result
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
