-- Drop existing tables if they exist to ensure a clean setup
-- DROP TABLE IF EXISTS environment_details;
-- DROP TABLE IF EXISTS environments;
-- DROP TABLE IF EXISTS products;
-- DROP TABLE IF EXISTS sections;
-- DROP TABLE IF EXISTS infra_types;
-- DROP TABLE IF EXISTS customers;
-- DROP TABLE IF EXISTS company;

-- Create tables

-- 1. Company table
CREATE TABLE IF NOT EXISTS company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Insert sample company data if it doesn't already exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM company WHERE name = 'My Company') THEN
        INSERT INTO company (name) VALUES ('My Company');
    END IF;
END $$;

-- 2. Infra types table
CREATE TABLE IF NOT EXISTS infra_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE -- INTERNAL, CUSTOMER
);

-- Insert sample infra types
INSERT INTO infra_types (name) VALUES ('INTERNAL') ON CONFLICT DO NOTHING;
INSERT INTO infra_types (name) VALUES ('CUSTOMER') ON CONFLICT DO NOTHING;

-- 3. Products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Insert sample products
INSERT INTO products (name) VALUES ('Product 1') ON CONFLICT DO NOTHING;
INSERT INTO products (name) VALUES ('Product 2') ON CONFLICT DO NOTHING;

-- 4. Customers table
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- 5. Sections table, linked to infra types
CREATE TABLE IF NOT EXISTS sections (
    id SERIAL PRIMARY KEY,
    infra_type_id INTEGER REFERENCES infra_types(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    environments TEXT[] -- e.g., DEV, QA, STAGING, etc.
);

-- Insert sample sections
INSERT INTO sections (infra_type_id, name, environments)
VALUES
    ((SELECT id FROM infra_types WHERE name = 'INTERNAL'), 'Product 1', ARRAY['DEV', 'QA', 'CONSULT', 'PRESALES']),
    ((SELECT id FROM infra_types WHERE name = 'INTERNAL'), 'Product 2', ARRAY['DEV', 'QA', 'STAGING']),
    ((SELECT id FROM infra_types WHERE name = 'CUSTOMER'), 'Vendor A', ARRAY['Product 1', 'Product 2']),
    ((SELECT id FROM infra_types WHERE name = 'CUSTOMER'), 'Vendor B', ARRAY['Product 1', 'Product 2'])
ON CONFLICT DO NOTHING;

-- 6. Environments table, linked to products
CREATE TABLE IF NOT EXISTS environments (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL
);

-- Insert sample environments
INSERT INTO environments (product_id, name)
VALUES
    ((SELECT id FROM products WHERE name = 'Product 1'), 'DEV'),
    ((SELECT id FROM products WHERE name = 'Product 1'), 'QA'),
    ((SELECT id FROM products WHERE name = 'Product 1'), 'CONSULT'),
    ((SELECT id FROM products WHERE name = 'Product 1'), 'PRESALES'),
    ((SELECT id FROM products WHERE name = 'Product 2'), 'DEV'),
    ((SELECT id FROM products WHERE name = 'Product 2'), 'QA'),
    ((SELECT id FROM products WHERE name = 'Product 2'), 'STAGING')
ON CONFLICT DO NOTHING;

-- 7. Environment details table, linked to environments
CREATE TABLE IF NOT EXISTS environment_details (
    id SERIAL PRIMARY KEY,
    environment_id INT NOT NULL REFERENCES environments(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    url VARCHAR(100),
    last_updated TIMESTAMP,
    status VARCHAR(20),
    contact VARCHAR(50),
    app_version VARCHAR(50),
    db_version VARCHAR(50),
    comments TEXT
);

-- Insert sample environment details for Product 1
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Dev', 'dev.example.com', '2021-08-19 21:30:00', 'Online', 'Samyak', 'develop-20240821.1', '7.2.0876', 'Testing this env so please check'),
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Smoke', 'smoke.example.com', '2021-08-19 21:30:00', 'Online', 'Samyak', 'develop-20240920.3', '7.2.0876', 'Testing this env so please check'),
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Stage', 'stage.example.com', '2021-08-19 21:30:00', 'Online', 'Samyak', 'develop-20240512.1', '7.2.0876', 'Testing this env so please check')
ON CONFLICT DO NOTHING;

-- Insert additional environment details for various setups

-- QA environments for Product 1
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Manual', 'manual.qa.example.com', '2021-08-23 08:00:00', 'Online', 'Alice', 'qa-manual-20240823.1', '7.2.0876', 'Manual QA environment for Product 1'),
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Auto', 'auto.qa.example.com', '2021-08-23 09:00:00', 'In Progress', 'Bob', 'qa-auto-20240901.2', '7.2.0877', 'Automated QA environment for Product 1'),
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Prelaunch', 'prelaunch.qa.example.com', '2021-08-23 10:00:00', 'Offline', 'Charlie', 'qa-prelaunch-20240915.1', '7.2.0878', 'Prelaunch QA environment for Product 1')
ON CONFLICT DO NOTHING;

-- Insert PRESALES environments for Product 1
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'PRESALES' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Demo', 'demo.presales.example.com', '2021-08-23 11:00:00', 'Online', 'David', 'presales-demo-20240901.1', '7.2.0880', 'Demo environment for Product 1 Presales'),
    ((SELECT id FROM environments WHERE name = 'PRESALES' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Sales', 'sales.presales.example.com', '2021-08-23 12:00:00', 'Offline', 'Eve', 'presales-sales-20240910.2', '7.2.0881', 'Sales environment for Product 1 Presales'),
    ((SELECT id FROM environments WHERE name = 'PRESALES' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Presales', 'presales.presales.example.com', '2021-08-23 13:00:00', 'In Progress', 'Frank', 'presales-presales-20240915.3', '7.2.0882', 'Presales environment for Product 1 Presales')
ON CONFLICT DO NOTHING;

-- Insert CONSULT environments for Product 1
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'CONSULT' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Tech', 'tech.consult.example.com', '2021-08-24 08:00:00', 'Online', 'Grace', 'consult-tech-20240905.1', '7.2.0879', 'Tech environment for Product 1 Consult'),
    ((SELECT id FROM environments WHERE name = 'CONSULT' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'SRE', 'sre.consult.example.com', '2021-08-24 09:00:00', 'Offline', 'Hank', 'consult-sre-20240912.2', '7.2.0883', 'SRE environment for Product 1 Consult')
ON CONFLICT DO NOTHING;

-- Add additional products, environments, and environment details similarly for a comprehensive test dataset.

-- Update statements for versioning environment details

-- Update Product 1 DEV environments
UPDATE environment_details
SET app_version = 'v1.0.0'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'Dev';

UPDATE environment_details
SET app_version = 'v1.0.1'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'Smoke';

UPDATE environment_details
SET app_version = 'v1.0.2'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'Stage';

-- Update Product 1 QA environments
UPDATE environment_details
SET app_version = 'v2.1.0'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'Manual';

UPDATE environment_details
SET app_version = 'v2.1.1'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'Auto';

UPDATE environment_details
SET app_version = 'v2.1.2'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'Prelaunch';

-- Update Product 1 PRESALES environments
UPDATE environment_details
SET app_version = 'v3.0.0'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'PRESALES' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'Demo';

UPDATE environment_details
SET app_version = 'v3.0.1'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'PRESALES' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'Sales';

UPDATE environment_details
SET app_version = 'v3.0.2'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'PRESALES' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'Presales';

-- Update Product 1 CONSULT environments
UPDATE environment_details
SET app_version = 'v4.2.0'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'CONSULT' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'Tech';

UPDATE environment_details
SET app_version = 'v4.2.1'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'CONSULT' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))
  AND name = 'SRE';

-- Update Product 2 QA environments
UPDATE environment_details
SET app_version = 'v5.0.0'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 2'))
  AND name = 'Manual';

UPDATE environment_details
SET app_version = 'v5.0.1'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 2'))
  AND name = 'Auto';

UPDATE environment_details
SET app_version = 'v5.0.2'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 2'))
  AND name = 'Prelaunch';

-- Update Product 2 STAGING environments
UPDATE environment_details
SET app_version = 'v6.1.0'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'STAGING' AND product_id = (SELECT id FROM products WHERE name = 'Product 2'))
  AND name = 'Release';

UPDATE environment_details
SET app_version = 'v6.1.1'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'STAGING' AND product_id = (SELECT id FROM products WHERE name = 'Product 2'))
  AND name = 'Launch';

UPDATE environment_details
SET app_version = 'v6.1.2'
WHERE environment_id = (SELECT id FROM environments WHERE name = 'STAGING' AND product_id = (SELECT id FROM products WHERE name = 'Product 2'))
  AND name = 'Hotfix';