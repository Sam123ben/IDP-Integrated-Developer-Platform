// src/types/index.ts
export type SystemStatus = 'operational' | 'maintenance' | 'degraded';

export interface SystemComponent {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  name: string; // Matches "name" from API
  type: string; // Matches "type" from API
  status: string; // Matches "status" from API
}

export interface MaintenanceUpdate {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  message: string;
}

export interface MaintenanceData {
  current_time: string;
  is_active: boolean;
  remaining_time_minutes: number;
  maintenance: {
    id: number;
    start_time: string; // Localized format from backend
    description: string;
    components: {
      maintenance: SystemComponent[];
      operational: SystemComponent[];
    };
    updates: {
      data: MaintenanceUpdate[];
      pagination: {
        current_page: number;
        per_page: number;
        total_updates: number;
      };
    };
  };
}