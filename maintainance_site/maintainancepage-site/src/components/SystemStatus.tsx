// src/components/SystemStatus.tsx
import React from 'react';
import { Server, Database, Globe } from 'lucide-react';
import type { SystemComponent } from '../types';

const getIcon = (type: string | undefined) => {
    switch (type?.toLowerCase()) {
        case 'api':
            return <Server className="w-5 h-5 text-gray-600" />;
        case 'database':
            return <Database className="w-5 h-5 text-gray-600" />;
        case 'cdn':
            return <Globe className="w-5 h-5 text-gray-600" />;
        default:
            return <Server className="w-5 h-5 text-gray-600" />;
    }
};

const statusColors = {
    operational: 'bg-green-100 text-green-700',
    maintenance: 'bg-yellow-100 text-yellow-700',
    degraded: 'bg-red-100 text-red-700'
} as const;

interface StatusItemProps {
    component: SystemComponent;
}

function StatusItem({ component }: StatusItemProps) {
    const status = component?.status?.toLowerCase() as keyof typeof statusColors;
    const colorClass = statusColors[status] || statusColors.operational;

    return (
        <div className="flex items-center justify-between p-3 bg-white rounded-lg shadow-sm">
            <div className="flex items-center space-x-3">
                {getIcon(component?.type)}
                <span className="font-medium">{component?.name || 'Unknown Component'}</span>
            </div>
            <span className={`px-3 py-1 rounded-full text-sm font-medium capitalize ${colorClass}`}>
                {component?.status?.toLowerCase() || 'unknown'}
            </span>
        </div>
    );
}

interface SystemStatusProps {
    components: SystemComponent[];
}

export function SystemStatus({ components }: SystemStatusProps) {
    if (!components?.length) {
        return (
            <div className="space-y-4">
                <h2 className="text-xl font-semibold text-gray-900">System Status</h2>
                <p className="text-gray-600">No system components to display</p>
            </div>
        );
    }

    return (
        <div className="space-y-4">
            <h2 className="text-xl font-semibold text-gray-900">System Status</h2>
            <div className="space-y-2">
                {components.map((component) => (
                    <StatusItem key={component.ID} component={component} />
                ))}
            </div>
        </div>
    );
}