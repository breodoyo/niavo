import { apiClient } from "../api/client";

export interface WorkItem {
  id: string;
  workflow_id: string;
  title: string;
  description?: string;
  status: string;
  priority: string;
  assigned_user_id?: string;
  due_date?: string;
  created_at: string;
  updated_at: string;
}

export interface CreateWorkItemRequest {
  workflow_id: string;
  title: string;
  description?: string;
  priority?: string;
  due_date?: string;
}

export interface UpdateWorkItemRequest {
  title: string;
  description?: string;
  status: string;
  priority: string;
  due_date?: string;
}

export async function listWorkItems(workflowId: string): Promise<WorkItem[]> {
  const response = await apiClient.get<WorkItem[]>("/workitems", {
    params: { workflow_id: workflowId },
  });
  return response.data;
}

export async function createWorkItem(data: CreateWorkItemRequest): Promise<WorkItem> {
  const response = await apiClient.post<WorkItem>("/workitems", data);
  return response.data;
}

export async function updateWorkItem(id: string, data: UpdateWorkItemRequest): Promise<WorkItem> {
  const response = await apiClient.patch<WorkItem>(`/workitems/${id}`, data);
  return response.data;
}

export async function deleteWorkItem(id: string): Promise<void> {
  await apiClient.delete(`/workitems/${id}`);
}