package models

// Guerabook model
type Guerabook struct {
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	UserID string `json:"userId,omitempty"`
}
