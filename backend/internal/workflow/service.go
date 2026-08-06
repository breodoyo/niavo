package workflow

import (
	"strings"
	"context"
)

type WorkflowService struct {
	repo *Repository
}

func NewService(repo *Repository) *WorkflowService {
	return &WorkflowService{repo: repo}
}

func (s *WorkflowService) CreateWorkflow(ctx context.Context, req CreateWorkflowRequest) (Workflow, error) {
	var description *string
	if strings.TrimSpace(req.Description) != "" {
		description = &req.Description
	}

	wf := Workflow{
		Name:        req.Name,
		Description: description,
	}

	return s.repo.CreateWorkflow(ctx, wf)
}

func (s *WorkflowService) ListWorkflows(ctx context.Context) ([]Workflow, error) {
	return s.repo.ListWorkflows(ctx)
}

func (s *WorkflowService) GetWorkflow(ctx context.Context, id string) (Workflow, error) {
	return s.repo.GetWorkflow(ctx, id)
}

func (s *WorkflowService) UpdateWorkflow(ctx context.Context, id string, req UpdateWorkflowRequest) (Workflow, error) {
	var description *string
	if strings.TrimSpace(req.Description) != "" {
		description = &req.Description
	}

	return s.repo.UpdateWorkflow(ctx, id, req.Name, description)
}
func (s *WorkflowService) DeleteWorkflow(ctx context.Context, id string) error {
	return s.repo.DeleteWorkflow(ctx, id)
}