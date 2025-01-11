// src/components/MaintenanceIcon.tsx
import React from 'react';
import { Settings } from 'lucide-react';

export function MaintenanceIcon() {
  return (
    <div className="flex justify-center">
      <Settings className="w-16 h-16 text-indigo-600 animate-spin-slow" />
    </div>
  );
}
