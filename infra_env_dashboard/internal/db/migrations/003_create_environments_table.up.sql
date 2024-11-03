-- 003_create_environments_table.up.sql
-- Create the environments table

CREATE TABLE IF NOT EXISTS environments (
    id SERIAL PRIMARY KEY,
    environment_name VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    environment_type VARCHAR(50) REFERENCES environment_types(name) ON DELETE SET NULL,
    environment_type_id INTEGER REFERENCES environment_types(id) ON DELETE SET NULL,
    group_id INTEGER REFERENCES groups(id) ON DELETE SET NULL,
    customer_id INTEGER REFERENCES groups(id) ON DELETE SET NULL,
    url VARCHAR(200),
    status VARCHAR(20),
    contact VARCHAR(100),
    app_version VARCHAR(50),
    db_version VARCHAR(50),
    comments TEXT
);

-- Sample insertions for environments with specific group associations
INSERT INTO environments (environment_name, description, environment_type, environment_type_id, group_id, customer_id, url, status, contact, app_version, db_version, comments)
VALUES 
    ('DEV', 'Development environment for internal testing', 
     'INTERNAL',  -- environment_type (name)
     (SELECT id FROM environment_types WHERE name = 'INTERNAL'), -- environment_type_id
     (SELECT id FROM groups WHERE name = 'dev'),  -- group_id
     (SELECT id FROM groups WHERE name = 'dev'),  -- customer_id
     'http://dev.example.com', 'Online', 'DevOps Team', 'v1.0.0', 'DB v10.5', 'No issues'),

    ('UAT', 'User acceptance testing environment for Customer 1', 
     'CUSTOMER',  -- environment_type (name)
     (SELECT id FROM environment_types WHERE name = 'CUSTOMER'), -- environment_type_id
     (SELECT id FROM groups WHERE name = 'Customer 1'),  -- group_id
     (SELECT id FROM groups WHERE name = 'Customer 1'),  -- customer_id
     'http://uat.customer1.com', 'Online', 'Customer 1 Support', 'v2.1.0', 'DB v11.3', 'In use by customer');