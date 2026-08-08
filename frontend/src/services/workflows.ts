import { apiClient } from "../api/client";

export interface Workflow {
  id: string;
  name: string;
  description?: string;
  created_at: string;
  updated_at: string;
}

export interface CreateWorkflowRequest {
  name: string;
  description?: string;
}

export interface UpdateWorkflowRequest {
  name: string;
  description?: string;
}

export async function listWorkflows(): Promise<Workflow[]> {
  const response = await apiClient.get<Workflow[]>("/workflows");
  return response.data;
}

export async function createWorkflow(data: CreateWorkflowRequest): Promise<Workflow> {
  const response = await apiClient.post<Workflow>("/workflows", data);
  return response.data;
}

export async function updateWorkflow(id: string, data: UpdateWorkflowRequest): Promise<Workflow> {
  const response = await apiClient.patch<Workflow>(`/workflows/${id}`, data);
  return response.data;
}

export async function deleteWorkflow(id: string): Promise<void> {
  await apiClient.delete(`/workflows/${id}`);
}