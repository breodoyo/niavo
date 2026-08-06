package workflow

import "context"

type Service interface {
	CreateWorkflow(ctx context.Context, req CreateWorkflowRequest) (Workflow, error)
	ListWorkflows(ctx context.Context) ([]Workflow, error)
	GetWorkflow(ctx context.Context, id string) (Workflow, error)
	UpdateWorkflow(ctx context.Context, id string, req UpdateWorkflowRequest) (Workflow, error)
	DeleteWorkflow(ctx context.Context, id string) error
}