-- Create groups table
-- 002_create_groups_table.up.sql

CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    environment_type_id INTEGER NOT NULL REFERENCES environment_types(id) ON DELETE CASCADE
);

-- Insert default groups for INTERNAL environments
INSERT INTO groups (name, environment_type_id)
VALUES 
    ('dev', (SELECT id FROM environment_types WHERE name = 'INTERNAL')),
    ('qa', (SELECT id FROM environment_types WHERE name = 'INTERNAL')),
    ('consultant', (SELECT id FROM environment_types WHERE name = 'INTERNAL')),
    ('presales', (SELECT id FROM environment_types WHERE name = 'INTERNAL')),
    ('release', (SELECT id FROM environment_types WHERE name = 'INTERNAL'))
ON CONFLICT (name) DO NOTHING;

-- Insert default groups for CUSTOMER environments
INSERT INTO groups (name, environment_type_id)
VALUES 
    ('Customer 1', (SELECT id FROM environment_types WHERE name = 'CUSTOMER')),
    ('Customer 2', (SELECT id FROM environment_types WHERE name = 'CUSTOMER')),
    ('Customer 3', (SELECT id FROM environment_types WHERE name = 'CUSTOMER')),
    ('Customer 4', (SELECT id FROM environment_types WHERE name = 'CUSTOMER')),
    ('Customer 5', (SELECT id FROM environment_types WHERE name = 'CUSTOMER'))
ON CONFLICT (name) DO NOTHING;