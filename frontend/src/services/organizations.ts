import { apiClient } from "../api/client";

export interface Organization {
  id: string;
  name: string;
  slug: string;
  created_at: string;
  updated_at: string;
}

export interface CreateOrganizationRequest {
  name: string;
  slug: string;
}

export async function listOrganizations(): Promise<Organization[]> {
  const response = await apiClient.get<Organization[]>("/organizations");
  return response.data;
}

export async function createOrganization(
  data: CreateOrganizationRequest
): Promise<Organization> {
  const response = await apiClient.post<Organization>("/organizations", data);
  return response.data;
}