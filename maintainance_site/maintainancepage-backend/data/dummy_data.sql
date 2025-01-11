-- dummy_data.sql

-- Insert a maintenance window
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM maintenance_windows WHERE description = 'We are performing scheduled system upgrades to improve platform stability and performance.'
    ) THEN
        INSERT INTO maintenance_windows (
            start_time,
            estimated_duration,
            description,
            created_at,
            updated_at
        ) VALUES (
            CURRENT_TIMESTAMP,
            120, -- 2 hours in minutes
            'We are performing scheduled system upgrades to improve platform stability and performance. Our team is working to minimize disruption.',
            CURRENT_TIMESTAMP,
            CURRENT_TIMESTAMP
        );
    END IF;
END $$;

-- Insert system components
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM system_components WHERE name = 'API Services' AND type = 'api'
    ) THEN
        INSERT INTO system_components (
            name,
            type,
            status,
            created_at,
            updated_at
        ) VALUES 
            ('API Services', 'api', 'maintenance', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM system_components WHERE name = 'Database' AND type = 'database'
    ) THEN
        INSERT INTO system_components (
            name,
            type,
            status,
            created_at,
            updated_at
        ) VALUES 
            ('Database', 'database', 'operational', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM system_components WHERE name = 'CDN' AND type = 'cdn'
    ) THEN
        INSERT INTO system_components (
            name,
            type,
            status,
            created_at,
            updated_at
        ) VALUES 
            ('CDN', 'cdn', 'operational', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
    END IF;
END $$;

-- Insert maintenance updates
DO $$
DECLARE
    maintenance_window_id INT;
BEGIN
    -- Get the maintenance window ID
    SELECT id INTO maintenance_window_id
    FROM maintenance_windows
    WHERE description = 'We are performing scheduled system upgrades to improve platform stability and performance.';

    IF maintenance_window_id IS NOT NULL THEN
        IF NOT EXISTS (
            SELECT 1 FROM maintenance_updates WHERE maintenance_window_id = maintenance_window_id AND message = 'Started system maintenance and security updates'
        ) THEN
            INSERT INTO maintenance_updates (
                maintenance_window_id,
                message,
                created_at
            ) VALUES 
                (maintenance_window_id, 'Started system maintenance and security updates', CURRENT_TIMESTAMP - interval '15 minutes');
        END IF;

        IF NOT EXISTS (
            SELECT 1 FROM maintenance_updates WHERE maintenance_window_id = maintenance_window_id AND message = 'Database backup completed successfully'
        ) THEN
            INSERT INTO maintenance_updates (
                maintenance_window_id,
                message,
                created_at
            ) VALUES 
                (maintenance_window_id, 'Database backup completed successfully', CURRENT_TIMESTAMP - interval '10 minutes');
        END IF;

        IF NOT EXISTS (
            SELECT 1 FROM maintenance_updates WHERE maintenance_window_id = maintenance_window_id AND message = 'Deploying system updates'
        ) THEN
            INSERT INTO maintenance_updates (
                maintenance_window_id,
                message,
                created_at
            ) VALUES 
                (maintenance_window_id, 'Deploying system updates', CURRENT_TIMESTAMP - interval '5 minutes');
        END IF;

        IF NOT EXISTS (
            SELECT 1 FROM maintenance_updates WHERE maintenance_window_id = maintenance_window_id AND message = 'Running final checks'
        ) THEN
            INSERT INTO maintenance_updates (
                maintenance_window_id,
                message,
                created_at
            ) VALUES 
                (maintenance_window_id, 'Running final checks', CURRENT_TIMESTAMP);
        END IF;
    END IF;
END $$;