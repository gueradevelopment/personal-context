package models

// Checklist model
type Checklist struct {
	ID             string `json:"id,omitempty"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	UserID         string `json:"userId"`
	IsTeamContext  bool   `json:"isTeamContext"`
	CompletionDate string `json:"completionDate,omitempty"`
	BoardID        string `json:"boardId,omitempty"`
}
