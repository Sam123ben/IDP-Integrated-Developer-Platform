// src/App.tsx
import React, { useEffect, useState } from 'react';
import { MaintenanceIcon } from './components/MaintenanceIcon';
import { StatusInfo } from './components/StatusInfo';
import { ContactInfo } from './components/ContactInfo';
import { CountdownTimer } from './components/CountdownTimer';
import { SystemStatus } from './components/SystemStatus';
import { UpdatesList } from './components/UpdatesList';
import type { MaintenanceData } from './types';
import { fetchMaintenanceData } from './services/api';

function App() {
    const [maintenanceData, setMaintenanceData] = useState<MaintenanceData | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<Error | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const data = await fetchMaintenanceData();
                console.log('Fetched data:', data);
                setMaintenanceData(data);
                setError(null);
            } catch (err) {
                console.error('Error fetching data:', err);
                setError(err instanceof Error ? err : new Error('Failed to fetch data'));
            } finally {
                setLoading(false);
            }
        };

        fetchData();
        const interval = setInterval(fetchData, 30000); // Refresh data every 30 seconds
        return () => clearInterval(interval);
    }, []);

    if (loading) {
        return (
            <div className="min-h-screen flex items-center justify-center bg-gray-50">
                <MaintenanceIcon />
                <p className="text-gray-600 mt-4">Loading maintenance status...</p>
            </div>
        );
    }

    if (error || !maintenanceData) {
        return (
            <div className="min-h-screen flex items-center justify-center bg-gray-50">
                <MaintenanceIcon />
                <h1 className="text-xl font-bold">Error loading maintenance status</h1>
                <p>{error?.message || 'Unknown error occurred'}</p>
            </div>
        );
    }

    const { maintenance } = maintenanceData;

    return (
        <div className="min-h-screen bg-gray-50">
            <div className="container mx-auto px-4 py-12">
                <div className="max-w-4xl mx-auto">
                    <div className="bg-white rounded-xl shadow-lg p-8 space-y-8">
                        {/* Header */}
                        <div className="text-center space-y-4">
                            <MaintenanceIcon />
                            <h1 className="text-3xl font-bold">Under Maintenance</h1>
                            <p className="text-gray-600">
                                We're currently performing scheduled maintenance to improve your
                                experience. Our team is working hard to get everything back online.
                            </p>
                        </div>

                        {/* Timer */}
                        <div className="text-center">
                            <CountdownTimer
                                startTime={new Date(maintenance.start_time).toISOString()} // Convert to ISO
                                estimatedDuration={maintenanceData.remaining_time_minutes}
                            />
                        </div>

                        {/* System Status & Updates */}
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
                            <div>
                                <SystemStatus
                                    components={[
                                        ...maintenance.components.maintenance,
                                        ...maintenance.components.operational,
                                    ]}
                                />
                            </div>
                            <div>
                                <UpdatesList updates={maintenance.updates.data} />
                            </div>
                        </div>

                        {/* Status Info */}
                        <div>
                            <StatusInfo
                                estimatedDuration={maintenanceData.remaining_time_minutes}
                                startTime={maintenance.start_time}
                                description={maintenance.description || 'No description provided'}
                            />
                        </div>

                        {/* Contact Info */}
                        <div>
                            <ContactInfo />
                        </div>

                        {/* Footer */}
                        <div className="border-t pt-6">
                            <p className="text-gray-600 text-center">
                                Follow our{' '}
                                <a
                                    href="/status"
                                    className="text-indigo-600 hover:text-indigo-700 font-medium"
                                >
                                    status page
                                </a>{' '}
                                for real-time updates
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default App;