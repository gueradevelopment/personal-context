package models

// Board model
type Board struct {
	ID            string `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	UserID        string `json:"userId,omitempty"`
	IsTeamContext bool   `json:"isTeamContext,omitempty"`
	GuerabookID   string `json:"guerabookId,omitempty"`
}
