import { apiClient } from "../api/client";
import type { User } from "./auth";

export async function listUsers(): Promise<User[]> {
  const response = await apiClient.get<User[]>("/users");
  return response.data;
}