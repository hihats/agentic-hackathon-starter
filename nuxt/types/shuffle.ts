export interface ShuffleRequest {
  participants: string[];
  group_size?: number;
  num_groups?: number;
}

export interface ShuffleResponse {
  groups: string[][];
}

export interface ErrorResponse {
  error: string;
  message: string;
}

export type GroupingMode = 'group_size' | 'num_groups';
