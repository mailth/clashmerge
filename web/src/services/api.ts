import { LinkConfig, MergeConfig } from '@/types';

class ApiService {
  private baseUrl: string;

  constructor() {
    this.baseUrl = "/api";
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`;

    const response = await fetch(url, {
      ...options,
      headers: {
        "Content-Type": "application/json",
        ...options.headers,
      },
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(`HTTP ${response.status}: ${errorText}`);
    }

    return response.json();
  }

  // Link Config APIs
  async getLinkConfigs(): Promise<LinkConfig[]> {
    return this.request<LinkConfig[]>("/link-configs");
  }

  async createLinkConfig(
    config: Omit<LinkConfig, "ID" | "CreatedAt" | "UpdatedAt">
  ): Promise<LinkConfig> {
    return this.request<LinkConfig>("/link-configs", {
      method: "POST",
      body: JSON.stringify(config),
    });
  }

  async updateLinkConfig(
    id: number,
    config: Omit<LinkConfig, "ID" | "CreatedAt" | "UpdatedAt">
  ): Promise<LinkConfig> {
    return this.request<LinkConfig>(`/link-configs/${id}`, {
      method: "PUT",
      body: JSON.stringify(config),
    });
  }

  async deleteLinkConfig(id: number): Promise<{ message: string }> {
    return this.request<{ message: string }>(`/link-configs/${id}`, {
      method: "DELETE",
    });
  }

  // Merge Config APIs
  async getMergeConfigs(): Promise<MergeConfig[]> {
    return this.request<MergeConfig[]>("/merge-configs");
  }

  async createMergeConfig(
    config: Omit<MergeConfig, "ID" | "CreatedAt" | "UpdatedAt">
  ): Promise<MergeConfig> {
    return this.request<MergeConfig>("/merge-configs", {
      method: "POST",
      body: JSON.stringify(config),
    });
  }

  async updateMergeConfig(
    id: number,
    config: Omit<MergeConfig, "ID" | "CreatedAt" | "UpdatedAt">
  ): Promise<MergeConfig> {
    return this.request<MergeConfig>(`/merge-configs/${id}`, {
      method: "PUT",
      body: JSON.stringify(config),
    });
  }

  async deleteMergeConfig(id: number): Promise<{ message: string }> {
    return this.request<{ message: string }>(`/merge-configs/${id}`, {
      method: "DELETE",
    });
  }
}

export const apiService = new ApiService();