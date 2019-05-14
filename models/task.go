package models

// Task model
type Task struct {
	ID              string `json:"id,omitempty"`
	Title           string `json:"title"`
	Description     string `json:"description,omitempty"`
	UserID          string `json:"userId"`
	IsTeamContext   bool   `json:"isTeamContext,omitempty"`
	CompletionState string `json:"completionState,omitempty"`
	CompletionDate  string `json:"completionDate,omitempty"`
	BoardID         string `json:"boardId,omitempty"`
	ChecklistID     string `json:"checklistId,omitempty"`
}
