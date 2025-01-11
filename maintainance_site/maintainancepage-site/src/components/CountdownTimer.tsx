// src/components/CountdownTimer.tsx
import React, { useState, useEffect } from 'react';
import { Timer } from 'lucide-react';
import { formatTime } from '../utils/timeUtils';

type CountdownTimerProps = {
    startTime: string; // ISO format
    estimatedDuration: number; // Minutes
};

export function CountdownTimer({ startTime, estimatedDuration }: CountdownTimerProps) {
    const [timeLeft, setTimeLeft] = useState<number>(0);

    useEffect(() => {
        const calculateTimeLeft = () => {
            const start = new Date(startTime).getTime();
            const end = start + estimatedDuration * 60 * 1000; // Minutes to milliseconds
            const now = Date.now();
            return Math.max(0, Math.floor((end - now) / 1000)); // Return seconds remaining
        };

        setTimeLeft(calculateTimeLeft());

        const timer = setInterval(() => {
            const remaining = calculateTimeLeft();
            setTimeLeft(remaining);
            if (remaining <= 0) {
                clearInterval(timer);
            }
        }, 1000);

        return () => clearInterval(timer);
    }, [startTime, estimatedDuration]);

    return (
        <div className="flex flex-col items-center space-y-3 bg-indigo-100 p-6 rounded-lg shadow-md">
            <Timer className="w-8 h-8 text-indigo-600" />
            <p className="text-indigo-600 font-medium">Estimated time remaining</p>
            <p className="text-indigo-800 text-3xl font-mono font-bold">{formatTime(timeLeft)}</p>
        </div>
    );
}