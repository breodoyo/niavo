package workitem

import (
	"context"
	"strings"
	"time"
)

type WorkItemService struct {
	repo *Repository
}

func NewService(repo *Repository) *WorkItemService {
	return &WorkItemService{repo: repo}
}

func ptrIfNotEmpty(s string) *string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	return &s
}

func parseDueDate(s string) (*time.Time, error) {
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *WorkItemService) CreateWorkItem(
	ctx context.Context,
	req CreateWorkItemRequest,
) (WorkItem, error) {
	priority := req.Priority
	if priority == "" {
		priority = "medium"
	}

	dueDate, err := parseDueDate(req.DueDate)
	if err != nil {
		return WorkItem{}, err
	}

	item := WorkItem{
		WorkflowID:     req.WorkflowID,
		Title:          req.Title,
		Description:    ptrIfNotEmpty(req.Description),
		Priority:       priority,
		AssignedUserID: ptrIfNotEmpty(req.AssignedUserID),
		DueDate:        dueDate,
	}

	return s.repo.CreateWorkItem(ctx, item)
}

func (s *WorkItemService) ListWorkItems(
	ctx context.Context,
	workflowID string,
) ([]WorkItem, error) {
	return s.repo.ListWorkItems(ctx, workflowID)
}

func (s *WorkItemService) GetWorkItem(
	ctx context.Context,
	id string,
) (WorkItem, error) {
	return s.repo.GetWorkItem(ctx, id)
}

func (s *WorkItemService) UpdateWorkItem(
	ctx context.Context,
	id string,
	req UpdateWorkItemRequest,
) (WorkItem, error) {
	status := req.Status
	if status == "" {
		status = "received"
	}

	priority := req.Priority
	if priority == "" {
		priority = "medium"
	}

	dueDate, err := parseDueDate(req.DueDate)
	if err != nil {
		return WorkItem{}, err
	}

	item := WorkItem{
		Title:          req.Title,
		Description:    ptrIfNotEmpty(req.Description),
		Status:         status,
		Priority:       priority,
		AssignedUserID: ptrIfNotEmpty(req.AssignedUserID),
		DueDate:        dueDate,
	}

	return s.repo.UpdateWorkItem(ctx, id, item)
}

func (s *WorkItemService) DeleteWorkItem(
	ctx context.Context,
	id string,
) error {
	return s.repo.DeleteWorkItem(ctx, id)
}