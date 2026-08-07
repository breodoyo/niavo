package workitem

import "context"

type Service interface {
	CreateWorkItem(ctx context.Context, req CreateWorkItemRequest) (WorkItem, error)
	ListWorkItems(ctx context.Context, workflowID string) ([]WorkItem, error)
	GetWorkItem(ctx context.Context, id string) (WorkItem, error)
	UpdateWorkItem(ctx context.Context, id string, req UpdateWorkItemRequest) (WorkItem, error)
	DeleteWorkItem(ctx context.Context, id string) error
}