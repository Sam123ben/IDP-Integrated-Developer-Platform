import React, { useEffect, useState } from 'react';
import { StatusInfo } from './StatusInfo';

const API_BASE_URL = 'http://localhost:8080/api';

export function ActiveMaintenance() {
  const [maintenanceData, setMaintenanceData] = useState<any>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchMaintenance = async () => {
      setLoading(true);
      setError(null);

      try {
        const response = await fetch(`${API_BASE_URL}/maintenance/active`);
        if (!response.ok) {
          throw new Error('Failed to fetch maintenance data');
        }
        const data = await response.json();
        setMaintenanceData(data.maintenance); // Save the "maintenance" object
      } catch (err: any) {
        setError(err.message || 'Something went wrong');
      } finally {
        setLoading(false);
      }
    };

    fetchMaintenance();
  }, []);

  if (loading) {
    return <p className="text-gray-700">Loading maintenance data...</p>;
  }

  if (error) {
    return <p className="text-red-600">Error: {error}</p>;
  }

  if (!maintenanceData) {
    return <p className="text-gray-700">No active maintenance window found.</p>;
  }

  // Render the StatusInfo component with fetched data
  return (
    <StatusInfo
      estimatedDuration={maintenanceData.estimated_duration}
      startTime={maintenanceData.start_time}
      description={maintenanceData.description}
    />
  );
}