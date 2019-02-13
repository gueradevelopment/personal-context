package models

// Board model
type Board struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	GuerabookID string `json:"guerabookId"`
}
