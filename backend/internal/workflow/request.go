package workflow

import (
	"errors"
	"strings"
)

type CreateWorkflowRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateWorkflowRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func (r CreateWorkflowRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("workflow name is required")
	}
	return nil
}

func (r UpdateWorkflowRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("workflow name is required")
	}
	return nil
}