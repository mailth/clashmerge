export interface LinkConfig {
  id?: number;
  name: string;
  clash_url: string;
  description: string;
  merge_config_id?: number;
  CreatedAt?: string;
  UpdatedAt?: string;
}

export interface MergeConfig {
  id?: number;
  name: string;
  rules?: string;
  proxies?: string;
  proxy_groups?: string;
  description: string;
  CreatedAt?: string;
  UpdatedAt?: string;
}

export interface ApiResponse<T> {
  data?: T;
  error?: string;
  message?: string;
}

export type SectionType = "link-config" | "merge-config";
