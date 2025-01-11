// app/types/form-schemas.ts
import { z } from 'zod';

export const baseSchema = z.object({
  email: z.string().email('Please enter a valid email address'),
  updateType: z.enum(['system_components', 'maintenance_windows', 'maintenance_updates']),
});

export const systemComponentsSchema = z.object({
  componentName: z.string().min(1, 'Component name is required'),
  componentType: z.string().min(1, 'Component type is required'),
  componentStatus: z.enum(['operational', 'degraded', 'maintenance']),
});

export const maintenanceWindowsSchema = baseSchema.extend({
  startTime: z.string().min(1, 'Required'),
  estimatedDuration: z.string().min(1, 'Required'),
  description: z.string().min(1, 'Required'),
});

export const maintenanceUpdatesSchema = baseSchema.extend({
  message: z.string().min(1, 'Required'),
});