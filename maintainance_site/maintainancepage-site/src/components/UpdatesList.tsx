// src/components/UpdatesList.tsx
import React from 'react';
import { CheckCircle2 } from 'lucide-react';
import type { MaintenanceUpdate } from '../types';

interface UpdatesListProps {
  updates: MaintenanceUpdate[];
}

export function UpdatesList({ updates }: UpdatesListProps) {
  const formatTime = (dateString: string) => {
    try {
      const date = new Date(dateString);
      return date.toLocaleTimeString('en-US', {
        hour: 'numeric',
        minute: '2-digit',
        hour12: true,
      });
    } catch (error) {
      console.error('Error formatting date:', error);
      return 'Invalid time';
    }
  };

  if (!updates?.length) {
    return (
      <div className="space-y-4">
        <h2 className="text-xl font-semibold text-gray-900">Recent Updates</h2>
        <p className="text-gray-600">No updates available</p>
      </div>
    );
  }

  // Sort updates by creation time, most recent first
  const sortedUpdates = [...updates].sort(
    (a, b) => new Date(b.CreatedAt).getTime() - new Date(a.CreatedAt).getTime()
  );

  return (
    <div className="space-y-4">
      <h2 className="text-xl font-semibold text-gray-900">Recent Updates</h2>
      <div className="space-y-3">
        {sortedUpdates.map((update) => (
          <div key={update.ID} className="flex items-start space-x-3">
            <CheckCircle2 className="w-5 h-5 text-green-500 mt-0.5" />
            <div>
              <span className="text-sm font-medium text-gray-500">
                {formatTime(update.CreatedAt)}
              </span>
              {/* Correctly access the "message" field */}
              <p className="text-gray-700">{update.message}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}