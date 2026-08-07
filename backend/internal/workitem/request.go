package workitem

import (
	"errors"
	"strings"
)

type CreateWorkItemRequest struct {
	WorkflowID     string  `json:"workflow_id"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Priority       string  `json:"priority"`
	AssignedUserID string  `json:"assigned_user_id"`
	DueDate        string  `json:"due_date"`
}

type UpdateWorkItemRequest struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	Priority       string `json:"priority"`
	AssignedUserID string `json:"assigned_user_id"`
	DueDate        string `json:"due_date"`
}

var validPriorities = map[string]bool{"low": true, "medium": true, "high": true, "urgent": true}
var validStatuses = map[string]bool{"received": true, "in_progress": true, "done": true}

func (r CreateWorkItemRequest) Validate() error {
	if strings.TrimSpace(r.WorkflowID) == "" {
		return errors.New("workflow_id is required")
	}
	if strings.TrimSpace(r.Title) == "" {
		return errors.New("title is required")
	}
	if r.Priority != "" && !validPriorities[r.Priority] {
		return errors.New("priority must be one of: low, medium, high, urgent")
	}
	return nil
}

func (r UpdateWorkItemRequest) Validate() error {
	if strings.TrimSpace(r.Title) == "" {
		return errors.New("title is required")
	}
	if r.Status != "" && !validStatuses[r.Status] {
		return errors.New("status must be one of: received, in_progress, done")
	}
	if r.Priority != "" && !validPriorities[r.Priority] {
		return errors.New("priority must be one of: low, medium, high, urgent")
	}
	return nil
}