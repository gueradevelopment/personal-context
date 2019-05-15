package models

// Checklist model
type Checklist struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	BoardID     string `json:"boardId"`
}
