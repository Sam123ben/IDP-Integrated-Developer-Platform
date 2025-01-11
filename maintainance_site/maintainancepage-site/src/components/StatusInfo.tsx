// src/components/StatusInfo.tsx
import React from 'react';
import { Clock, Calendar, AlertCircle } from 'lucide-react';

interface StatusInfoProps {
  estimatedDuration: number;
  startTime: string;
  description: string; // Fetch dynamically
}

export function StatusInfo({ estimatedDuration, startTime, description }: StatusInfoProps) {
  const formatDuration = (minutes: number): string => {
    const hours = Math.floor(minutes / 60);
    const remainingMinutes = minutes % 60;
    if (hours === 0) {
      return `${minutes} minutes`;
    }
    return `${hours} hour${hours > 1 ? 's' : ''}${remainingMinutes > 0 ? ` ${remainingMinutes} minutes` : ''}`;
  };

  const formatStartTime = (dateString: string): string => {
    try {
      const date = new Date(dateString);
      return date.toLocaleString('en-US', {
        weekday: 'long',
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: 'numeric',
        minute: '2-digit',
        hour12: true
      });
    } catch (error) {
      console.error('Error formatting date:', error);
      return 'Invalid Date';
    }
  };

  return (
    <div className="bg-indigo-50 rounded-xl p-6 space-y-4">
      <div className="flex items-center space-x-3">
        <Clock className="w-5 h-5 text-indigo-600" />
        <p className="text-gray-700">
          <span className="font-semibold">Estimated downtime:</span>{' '}
          {formatDuration(estimatedDuration)}
        </p>
      </div>
      <div className="flex items-center space-x-3">
        <Calendar className="w-5 h-5 text-indigo-600" />
        <p className="text-gray-700">
          <span className="font-semibold">Started at:</span>{' '}
          {formatStartTime(startTime)}
        </p>
      </div>
      <div className="flex items-center space-x-3">
        <AlertCircle className="w-5 h-5 text-indigo-600" />
        <p className="text-gray-700">
          <span className="font-semibold">Status:</span>{' '}
          {description || 'No description available'}
        </p>
      </div>
    </div>
  );
}