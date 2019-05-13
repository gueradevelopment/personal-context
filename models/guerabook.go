package models

// Guerabook model
type Guerabook struct {
	ID     string `json:"id,omitempty"`
	Title  string `json:"title"`
	UserID string `json:"userId"`
}
