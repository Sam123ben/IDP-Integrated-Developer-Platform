// src/utils/timeUtils.ts
export function formatTime(timeInSeconds: number): string {
  if (timeInSeconds <= 0) return "00:00:00";
  
  const hours = Math.floor(timeInSeconds / 3600);
  const minutes = Math.floor((timeInSeconds % 3600) / 60);
  const seconds = timeInSeconds % 60;

  return [hours, minutes, seconds]
    .map(v => v.toString().padStart(2, '0'))
    .join(':');
}

export function isValidDate(dateString: string): boolean {
  const date = new Date(dateString);
  return date instanceof Date && !isNaN(date.getTime());
}

export function formatDateTime(dateString: string): string {
  try {
    const date = new Date(dateString);
    return date.toLocaleTimeString('en-US', {
      hour: 'numeric',
      minute: '2-digit',
      hour12: true
    });
  } catch {
    return 'Invalid Time';
  }
}