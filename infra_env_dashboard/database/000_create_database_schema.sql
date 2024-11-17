-- Drop existing tables if they exist to ensure a clean setup
DROP TABLE IF EXISTS environment_details;
DROP TABLE IF EXISTS environments;
DROP TABLE IF EXISTS customer_products;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS company;

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

-- 2. Products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Insert sample products
INSERT INTO products (name)
VALUES ('Product 1'), ('Product 2')
ON CONFLICT DO NOTHING;

-- 3. Customers table
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Insert sample customers
INSERT INTO customers (name)
VALUES ('Vendor A'), ('Vendor B')
ON CONFLICT DO NOTHING;

-- 4. Customer Products join table
CREATE TABLE IF NOT EXISTS customer_products (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id) ON DELETE CASCADE
);

-- Insert customer-product relationships
INSERT INTO customer_products (customer_id, product_id)
VALUES
    ((SELECT id FROM customers WHERE name = 'Vendor A'), (SELECT id FROM products WHERE name = 'Product 1')),
    ((SELECT id FROM customers WHERE name = 'Vendor A'), (SELECT id FROM products WHERE name = 'Product 2')),
    ((SELECT id FROM customers WHERE name = 'Vendor B'), (SELECT id FROM products WHERE name = 'Product 1')),
    ((SELECT id FROM customers WHERE name = 'Vendor B'), (SELECT id FROM products WHERE name = 'Product 2'))
ON CONFLICT DO NOTHING;

-- 5. Environments table, linked to products and optionally customers
CREATE TABLE IF NOT EXISTS environments (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL
);

-- Insert internal environments
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

-- Insert customer environments
-- Vendor A Environments for Product 1
INSERT INTO environments (product_id, customer_id, name)
VALUES
    ((SELECT id FROM products WHERE name = 'Product 1'), (SELECT id FROM customers WHERE name = 'Vendor A'), 'UAT'),
    ((SELECT id FROM products WHERE name = 'Product 1'), (SELECT id FROM customers WHERE name = 'Vendor A'), 'TRAIN'),
    ((SELECT id FROM products WHERE name = 'Product 1'), (SELECT id FROM customers WHERE name = 'Vendor A'), 'SUPPORT'),
    ((SELECT id FROM products WHERE name = 'Product 1'), (SELECT id FROM customers WHERE name = 'Vendor A'), 'PROD'),

-- Vendor A Environments for Product 2
    ((SELECT id FROM products WHERE name = 'Product 2'), (SELECT id FROM customers WHERE name = 'Vendor A'), 'UAT01'),
    ((SELECT id FROM products WHERE name = 'Product 2'), (SELECT id FROM customers WHERE name = 'Vendor A'), 'UAT02'),
    ((SELECT id FROM products WHERE name = 'Product 2'), (SELECT id FROM customers WHERE name = 'Vendor A'), 'PREPROD'),
    ((SELECT id FROM products WHERE name = 'Product 2'), (SELECT id FROM customers WHERE name = 'Vendor A'), 'PROD'),

-- Vendor B Environments for Product 1
    ((SELECT id FROM products WHERE name = 'Product 1'), (SELECT id FROM customers WHERE name = 'Vendor B'), 'UAT'),
    ((SELECT id FROM products WHERE name = 'Product 1'), (SELECT id FROM customers WHERE name = 'Vendor B'), 'SIT'),
    ((SELECT id FROM products WHERE name = 'Product 1'), (SELECT id FROM customers WHERE name = 'Vendor B'), 'PROD-Blue'),
    ((SELECT id FROM products WHERE name = 'Product 1'), (SELECT id FROM customers WHERE name = 'Vendor B'), 'PROD-Green'),

-- Vendor B Environments for Product 2
    ((SELECT id FROM products WHERE name = 'Product 2'), (SELECT id FROM customers WHERE name = 'Vendor B'), 'UAT'),
    ((SELECT id FROM products WHERE name = 'Product 2'), (SELECT id FROM customers WHERE name = 'Vendor B'), 'SUPPORT'),
    ((SELECT id FROM products WHERE name = 'Product 2'), (SELECT id FROM customers WHERE name = 'Vendor B'), 'PROD-Blue'),
    ((SELECT id FROM products WHERE name = 'Product 2'), (SELECT id FROM customers WHERE name = 'Vendor B'), 'PROD-Green')
ON CONFLICT DO NOTHING;

-- 6. Environment details table, linked to environments
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

-- Insert sample environment details for internal environments

-- DEV environments for Product 1
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'Dev', 'dev.example.com', '2021-08-19 21:30:00', 'Online', 'Samyak', 'v1.0.0', '7.2.0876', 'Testing this env so please check'),
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'Smoke', 'smoke.example.com', '2021-08-20 21:30:00', 'Online', 'Samyak', 'v1.0.1', '7.2.0876', 'Smoke testing environment'),
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'Stage', 'stage.example.com', '2021-08-21 21:30:00', 'Online', 'Samyak', 'v1.0.2', '7.2.0876', 'Staging environment for final checks')
ON CONFLICT DO NOTHING;

-- QA environments for Product 1
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'Manual', 'manual.qa.example.com', '2021-08-23 08:00:00', 'Online', 'Alice', 'v2.1.0', '7.2.0876', 'Manual QA environment for Product 1'),
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'Auto', 'auto.qa.example.com', '2021-08-24 09:00:00', 'In Progress', 'Bob', 'v2.1.1', '7.2.0877', 'Automated QA environment for Product 1'),
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'Prelaunch', 'prelaunch.qa.example.com', '2021-08-25 10:00:00', 'Offline', 'Charlie', 'v2.1.2', '7.2.0878', 'Prelaunch QA environment for Product 1')
ON CONFLICT DO NOTHING;

-- PRESALES environments for Product 1
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'PRESALES' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'Demo', 'demo.presales.example.com', '2021-08-26 11:00:00', 'Online', 'David', 'v3.0.0', '7.2.0880', 'Demo environment for Product 1 Presales'),
    ((SELECT id FROM environments WHERE name = 'PRESALES' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'Sales', 'sales.presales.example.com', '2021-08-27 12:00:00', 'Offline', 'Eve', 'v3.0.1', '7.2.0881', 'Sales environment for Product 1 Presales'),
    ((SELECT id FROM environments WHERE name = 'PRESALES' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'Presales', 'presales.presales.example.com', '2021-08-28 13:00:00', 'In Progress', 'Frank', 'v3.0.2', '7.2.0882', 'Presales environment for Product 1 Presales')
ON CONFLICT DO NOTHING;

-- CONSULT environments for Product 1
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'CONSULT' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'Tech', 'tech.consult.example.com', '2021-08-29 08:00:00', 'Online', 'Grace', 'v4.2.0', '7.2.0879', 'Tech environment for Product 1 Consult'),
    ((SELECT id FROM environments WHERE name = 'CONSULT' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id IS NULL), 'SRE', 'sre.consult.example.com', '2021-08-30 09:00:00', 'Offline', 'Hank', 'v4.2.1', '7.2.0883', 'SRE environment for Product 1 Consult')
ON CONFLICT DO NOTHING;

-- DEV environments for Product 2
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id IS NULL), 'Dev', 'dev.product2.example.com', '2021-09-01 10:00:00', 'Online', 'Ivan', 'v5.0.0', '7.2.0876', 'Development environment for Product 2'),
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id IS NULL), 'Integration', 'integration.product2.example.com', '2021-09-02 11:00:00', 'Online', 'Ivan', 'v5.0.1', '7.2.0876', 'Integration environment for Product 2')
ON CONFLICT DO NOTHING;

-- QA environments for Product 2
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id IS NULL), 'Manual', 'manual.qa.product2.example.com', '2021-09-03 08:00:00', 'Online', 'Julia', 'v5.1.0', '7.2.0876', 'Manual QA environment for Product 2'),
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id IS NULL), 'Auto', 'auto.qa.product2.example.com', '2021-09-04 09:00:00', 'In Progress', 'Kevin', 'v5.1.1', '7.2.0877', 'Automated QA environment for Product 2'),
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id IS NULL), 'Prelaunch', 'prelaunch.qa.product2.example.com', '2021-09-05 10:00:00', 'Offline', 'Laura', 'v5.1.2', '7.2.0878', 'Prelaunch QA environment for Product 2')
ON CONFLICT DO NOTHING;

-- STAGING environments for Product 2
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'STAGING' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id IS NULL), 'Release', 'release.staging.product2.example.com', '2021-09-06 11:00:00', 'Online', 'Michael', 'v6.1.0', '7.2.0880', 'Release staging environment for Product 2'),
    ((SELECT id FROM environments WHERE name = 'STAGING' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id IS NULL), 'Launch', 'launch.staging.product2.example.com', '2021-09-07 12:00:00', 'Offline', 'Nina', 'v6.1.1', '7.2.0881', 'Launch staging environment for Product 2'),
    ((SELECT id FROM environments WHERE name = 'STAGING' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id IS NULL), 'Hotfix', 'hotfix.staging.product2.example.com', '2021-09-08 13:00:00', 'In Progress', 'Oscar', 'v6.1.2', '7.2.0882', 'Hotfix staging environment for Product 2')
ON CONFLICT DO NOTHING;

-- Insert sample environment details for customer environments

-- Vendor A Environments for Product 1
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    (
        (SELECT id FROM environments WHERE name = 'UAT' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')),
        'UAT',
        'uat.vendorA.product1.example.com',
        '2024-11-16 10:00:00',
        'Online',
        'Derrick',
        'v1.0.0',
        '7.2.0876',
        'UAT Environment for Vendor A Product 1'
    ),
    (
        (SELECT id FROM environments WHERE name = 'TRAIN' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')),
        'TRAIN',
        'train.vendorA.product1.example.com',
        '2024-11-16 11:00:00',
        'Online',
        'Derrick',
        'v1.0.1',
        '7.2.0876',
        'Training Environment for Vendor A Product 1'
    ),
    (
        (SELECT id FROM environments WHERE name = 'SUPPORT' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')),
        'SUPPORT',
        'support.vendorA.product1.example.com',
        '2024-11-16 12:00:00',
        'Online',
        'Derrick',
        'v1.0.2',
        '7.2.0876',
        'Support Environment for Vendor A Product 1'
    ),
    (
        (SELECT id FROM environments WHERE name = 'PROD' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')),
        'PROD',
        'prod.vendorA.product1.example.com',
        '2024-11-16 13:00:00',
        'Online',
        'Derrick',
        'v1.1.0',
        '7.2.0876',
        'Production Environment for Vendor A Product 1'
    )
ON CONFLICT DO NOTHING;

-- Vendor A Environments for Product 2
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    (
        (SELECT id FROM environments WHERE name = 'UAT01' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')),
        'UAT01',
        'uat01.vendorA.product2.example.com',
        '2024-11-16 10:00:00',
        'Online',
        'Mahesh',
        'v2.0.0',
        '7.2.0876',
        'UAT01 Environment for Vendor A Product 2'
    ),
    (
        (SELECT id FROM environments WHERE name = 'UAT02' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')),
        'UAT02',
        'uat02.vendorA.product2.example.com',
        '2024-11-16 11:00:00',
        'Online',
        'Mahesh',
        'v2.0.1',
        '7.2.0876',
        'UAT02 Environment for Vendor A Product 2'
    ),
    (
        (SELECT id FROM environments WHERE name = 'PREPROD' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')),
        'PREPROD',
        'preprod.vendorA.product2.example.com',
        '2024-11-16 12:00:00',
        'Online',
        'Mahesh',
        'v2.0.2',
        '7.2.0876',
        'Preprod Environment for Vendor A Product 2'
    ),
    (
        (SELECT id FROM environments WHERE name = 'PROD' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')),
        'PROD',
        'prod.vendorA.product2.example.com',
        '2024-11-16 13:00:00',
        'Online',
        'Mahesh',
        'v2.1.0',
        '7.2.0876',
        'Production Environment for Vendor A Product 2'
    )
ON CONFLICT DO NOTHING;

-- Vendor B Environments for Product 1
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    (
        (SELECT id FROM environments WHERE name = 'UAT' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor B')),
        'UAT',
        'uat.vendorB.product1.example.com',
        '2024-11-16 10:00:00',
        'Online',
        'Ranjeet',
        'v3.0.0',
        '7.2.0876',
        'UAT Environment for Vendor B Product 1'
    ),
    (
        (SELECT id FROM environments WHERE name = 'SIT' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor B')),
        'SIT',
        'sit.vendorB.product1.example.com',
        '2024-11-16 11:00:00',
        'Online',
        'Ranjeet',
        'v3.0.1',
        '7.2.0876',
        'SIT Environment for Vendor B Product 1'
    ),
    (
        (SELECT id FROM environments WHERE name = 'PROD-Blue' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor B')),
        'PROD-Blue',
        'prod-blue.vendorB.product1.example.com',
        '2024-11-16 12:00:00',
        'Online',
        'Ranjeet',
        'v3.1.0',
        '7.2.0876',
        'Production Blue Environment for Vendor B Product 1'
    ),
    (
        (SELECT id FROM environments WHERE name = 'PROD-Green' AND product_id = (SELECT id FROM products WHERE name = 'Product 1') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor B')),
        'PROD-Green',
        'prod-green.vendorB.product1.example.com',
        '2024-11-16 13:00:00',
        'Online',
        'Ranjeet',
        'v3.1.1',
        '7.2.0876',
        'Production Green Environment for Vendor B Product 1'
    )
ON CONFLICT DO NOTHING;

-- Vendor B Environments for Product 2
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    (
        (SELECT id FROM environments WHERE name = 'UAT' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor B')),
        'UAT',
        'uat.vendorB.product2.example.com',
        '2024-11-16 10:00:00',
        'Online',
        'Danny',
        'v4.0.0',
        '7.2.0876',
        'UAT Environment for Vendor B Product 2'
    ),
    (
        (SELECT id FROM environments WHERE name = 'SUPPORT' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor B')),
        'SUPPORT',
        'support.vendorB.product2.example.com',
        '2024-11-16 11:00:00',
        'Online',
        'Danny',
        'v4.0.1',
        '7.2.0876',
        'Support Environment for Vendor B Product 2'
    ),
    (
        (SELECT id FROM environments WHERE name = 'PROD-Blue' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor B')),
        'PROD-Blue',
        'prod-blue.vendorB.product2.example.com',
        '2024-11-16 12:00:00',
        'Online',
        'Danny',
        'v4.1.0',
        '7.2.0876',
        'Production Blue Environment for Vendor B Product 2'
    ),
    (
        (SELECT id FROM environments WHERE name = 'PROD-Green' AND product_id = (SELECT id FROM products WHERE name = 'Product 2') AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor B')),
        'PROD-Green',
        'prod-green.vendorB.product2.example.com',
        '2024-11-16 13:00:00',
        'Online',
        'Danny',
        'v4.1.1',
        '7.2.0876',
        'Production Green Environment for Vendor B Product 2'
    )
ON CONFLICT DO NOTHING;