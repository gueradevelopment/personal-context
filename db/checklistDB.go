package db

import (
	"errors"

	"github.com/gueradevelopment/personal-context/models"
)

// ChecklistDB - Checklist model database accessor
type ChecklistDB struct{}

var (
	checklistItems = make(map[string]models.Checklist)
)

// Get - retrieves a single resource
func (db *ChecklistDB) Get(id string, c chan Result) {
	defer close(c)
	result := Result{}
	for ID, item := range checklistItems {
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
func (db *ChecklistDB) GetAll(c chan ResultArray, where map[string][]string) {
	defer close(c)
	result := ResultArray{}
	var arr = []Model{}
	var boardID string
	if where["boardId"] != nil {
		boardID = where["boardId"][0]
	}
	for _, v := range checklistItems {
		if boardID != "" && v.BoardID == boardID {
			arr = append(arr, v)
		}
		if boardID == "" {
			arr = append(arr, v)
		}
	}
	result.Result = arr
	c <- result
}

// Add - creates a resource
func (db *ChecklistDB) Add(item models.Checklist, c chan Result) {
	defer close(c)
	result := Result{}
	if checklistItems[item.ID] == (models.Checklist{}) {
		checklistItems[item.ID] = item
		result.Result = item
	} else {
		result.Err = errors.New("Duplicated ID")
	}
	c <- result
}

// Edit - updates a resource
func (db *ChecklistDB) Edit(item models.Checklist, c chan Result) {
	defer close(c)
	result := Result{}
	if checklistItems[item.ID] == (models.Checklist{}) {
		result.Err = errors.New("No such ID")
	} else {
		checklistItems[item.ID] = item
		result.Result = item
	}
	c <- result
}

// Delete - deletes a resource
func (db *ChecklistDB) Delete(id string, c chan Result) {
	defer close(c)
	result := Result{}
	item := checklistItems[id]
	if item == (models.Checklist{}) {
		result.Err = errors.New("No such ID")
	} else {
		result.Result = item
		delete(checklistItems, id)
	}
	c <- result
}
