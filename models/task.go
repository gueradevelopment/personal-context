package models

// Task model
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	BoardID     string `json:"boardId"`
	ChecklistID string `json:"checklistId"`
}
