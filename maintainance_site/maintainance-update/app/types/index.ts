// app/types/index.ts
export interface ComponentStatus {
  key: string;
  value: string;
}

export interface SystemComponent {
  id: number;
  name: string;
  type: string;
  status: ComponentStatus;
}

export interface MaintenanceWindow {
  start_time: Date;
  estimated_duration: string;
  description: string;
}

export interface MaintenanceUpdate {
  message: string;
}

export interface FormData {
  email: string;
  updateType: 'system_components' | 'maintenance_windows' | 'maintenance_updates';
  systemComponent?: SystemComponent;
  maintenanceWindow?: MaintenanceWindow;
  maintenanceUpdate?: MaintenanceUpdate;
}