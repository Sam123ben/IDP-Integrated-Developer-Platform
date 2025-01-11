// src/services/api.ts
import type { MaintenanceData } from '../types';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api';

export async function fetchMaintenanceData(): Promise<MaintenanceData> {
    const response = await fetch(`${API_BASE_URL}/maintenance/active`);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
}
