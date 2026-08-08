package workitem

import "time"

type WorkItem struct {
	ID             string     `json:"id"`
	WorkflowID     string     `json:"workflow_id"`
	Title          string     `json:"title"`
	Description    *string    `json:"description,omitempty"`
	Status         string     `json:"status"`
	Priority       string     `json:"priority"`
	AssignedUserID *string    `json:"assigned_user_id,omitempty"`
	DueDate        *time.Time `json:"due_date,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}