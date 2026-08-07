package workflow

import "context"

type MockWorkflowService struct{}

func (m *MockWorkflowService) CreateWorkflow(ctx context.Context, req CreateWorkflowRequest) (Workflow, error) {
	desc := req.Description
	return Workflow{
		ID:          "123",
		Name:        req.Name,
		Description: &desc,
	}, nil
}

func (m *MockWorkflowService) ListWorkflows(ctx context.Context) ([]Workflow, error) {
	return []Workflow{}, nil
}

func (m *MockWorkflowService) GetWorkflow(ctx context.Context, id string) (Workflow, error) {
	return Workflow{ID: id, Name: "Mock Workflow"}, nil
}

func (m *MockWorkflowService) UpdateWorkflow(ctx context.Context, id string, req UpdateWorkflowRequest) (Workflow, error) {
	desc := req.Description
	return Workflow{
		ID:          id,
		Name:        req.Name,
		Description: &desc,
	}, nil
}

func (m *MockWorkflowService) DeleteWorkflow(ctx context.Context, id string) error {
	return nil
}