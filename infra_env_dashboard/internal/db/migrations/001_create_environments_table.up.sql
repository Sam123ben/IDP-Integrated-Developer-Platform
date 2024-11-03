-- 001_create_environments_table.up.sql

-- Check if the table 'environments' exists
DO $$
BEGIN
    -- Create table if it does not exist
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'environments') THEN
        CREATE TABLE environments (
            id SERIAL PRIMARY KEY,
            environment_name VARCHAR(50) NOT NULL,
            description TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            environment_type VARCHAR(20) CHECK (environment_type IN ('INTERNAL', 'EXTERNAL')),
            group_name VARCHAR(50),
            customer_name VARCHAR(100),
            url VARCHAR(200),
            status VARCHAR(20),
            contact VARCHAR(100),
            app_version VARCHAR(50),
            db_version VARCHAR(50),
            comments TEXT
        );
    ELSE
        -- If the table exists, alter the table to ensure all columns match the desired schema

        -- Add missing columns if they don't exist
        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'environment_name') THEN
            ALTER TABLE environments ADD COLUMN environment_name VARCHAR(50) NOT NULL;
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'description') THEN
            ALTER TABLE environments ADD COLUMN description TEXT NOT NULL;
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'created_at') THEN
            ALTER TABLE environments ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'updated_at') THEN
            ALTER TABLE environments ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'environment_type') THEN
            ALTER TABLE environments ADD COLUMN environment_type VARCHAR(20) CHECK (environment_type IN ('INTERNAL', 'EXTERNAL'));
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'group_name') THEN
            ALTER TABLE environments ADD COLUMN group_name VARCHAR(50);
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'customer_name') THEN
            ALTER TABLE environments ADD COLUMN customer_name VARCHAR(100);
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'url') THEN
            ALTER TABLE environments ADD COLUMN url VARCHAR(200);
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'status') THEN
            ALTER TABLE environments ADD COLUMN status VARCHAR(20);
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'contact') THEN
            ALTER TABLE environments ADD COLUMN contact VARCHAR(100);
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'app_version') THEN
            ALTER TABLE environments ADD COLUMN app_version VARCHAR(50);
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'db_version') THEN
            ALTER TABLE environments ADD COLUMN db_version VARCHAR(50);
        END IF;

        IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'environments' AND column_name = 'comments') THEN
            ALTER TABLE environments ADD COLUMN comments TEXT;
        END IF;

        -- Update constraints or defaults if necessary

        -- Ensure CHECK constraint on environment_type
        BEGIN
            ALTER TABLE environments DROP CONSTRAINT IF EXISTS environment_type_check;
            ALTER TABLE environments ADD CONSTRAINT environment_type_check CHECK (environment_type IN ('INTERNAL', 'EXTERNAL'));
        EXCEPTION
            WHEN duplicate_object THEN NULL; -- Ignore if constraint already exists
        END;
        
        -- Ensure default values for created_at and updated_at columns
        ALTER TABLE environments ALTER COLUMN created_at SET DEFAULT CURRENT_TIMESTAMP;
        ALTER TABLE environments ALTER COLUMN updated_at SET DEFAULT CURRENT_TIMESTAMP;

    END IF;
END $$;