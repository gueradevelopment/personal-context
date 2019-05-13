package models

// Board model
type Board struct {
	ID            string `json:"id,omitempty"`
	Title         string `json:"title"`
	UserID        string `json:"userId"`
	IsTeamContext bool   `json:"isTeamContext"`
	GuerabookID   string `json:"guerabookId,omitempty"`
}
