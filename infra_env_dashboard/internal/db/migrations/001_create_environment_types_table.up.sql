-- Create the environment_types table
-- 001_create_environment_types_table.up.sql

CREATE TABLE IF NOT EXISTS environment_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL -- 'INTERNAL' or 'CUSTOMER'
);

-- Insert default environment types
INSERT INTO environment_types (name) VALUES ('INTERNAL'), ('CUSTOMER')
ON CONFLICT (name) DO NOTHING;