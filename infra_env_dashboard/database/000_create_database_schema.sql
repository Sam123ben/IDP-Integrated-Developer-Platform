-- 1. Drop existing data tables if necessary (optional)
-- DROP TABLE IF EXISTS customer_environment_groups;
-- DROP TABLE IF EXISTS product_environment_groups;
-- DROP TABLE IF EXISTS infra_types;
-- DROP TABLE IF EXISTS sections;
-- DROP TABLE IF EXISTS env_groups;
-- DROP TABLE IF EXISTS customer_products;
-- DROP TABLE IF EXISTS customers;
-- DROP TABLE IF EXISTS applications;
-- DROP TABLE IF EXISTS environments;
-- DROP TABLE IF EXISTS products;
-- DROP TABLE IF EXISTS company;

-- 0. Create the company table
CREATE TABLE IF NOT EXISTS company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Insert company name if it does not exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM company WHERE name = 'My Company') THEN
        INSERT INTO company (name) VALUES ('My Company');
    END IF;
END $$;

-- 1. Infra Types table: Holds the types like INTERNAL and CUSTOMER
CREATE TABLE IF NOT EXISTS infra_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE -- e.g., INTERNAL, CUSTOMER
);

-- 2. Sections table: Represents sections under each infra type, such as Product 1, Vendor A, etc.
CREATE TABLE IF NOT EXISTS sections (
    id SERIAL PRIMARY KEY,
    infra_type_id INT REFERENCES infra_types(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL, -- e.g., Product 1, Vendor A, etc.
    UNIQUE (infra_type_id, name) -- Each section is unique within an infra type
);

-- 3. Environment Groups table: Represents environment group names (DEV, QA, CONSULT) under each section for INTERNAL type
CREATE TABLE IF NOT EXISTS environment_groups (
    id SERIAL PRIMARY KEY,
    section_id INT REFERENCES sections(id) ON DELETE CASCADE, -- Linked to the section under INTERNAL type
    name VARCHAR(50) NOT NULL, -- e.g., DEV, QA, CONSULT, etc.
    UNIQUE (section_id, name) -- Unique environment group name within a section
);

-- 4. Environments table: Stores specific environments within each environment group
CREATE TABLE IF NOT EXISTS environments (
    id SERIAL PRIMARY KEY,
    env_group_id INT REFERENCES environment_groups(id) ON DELETE CASCADE, -- Linked to an environment group
    name VARCHAR(50) NOT NULL, -- Environment name, e.g., DEV, QA, STAGING
    UNIQUE (env_group_id, name) -- Unique environment name within an environment group
);

-- 5. Customers table: Represents customers like Vendor A, Vendor B under CUSTOMER type
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE -- e.g., Vendor A, Vendor B
);

-- 6. Customer Products table: Links customers to products under CUSTOMER infra type
CREATE TABLE IF NOT EXISTS customer_products (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
    product_name VARCHAR(50) NOT NULL, -- Product name associated with the customer
    UNIQUE (customer_id, product_name) -- Ensures each customer has unique products
);

-- 7. Insert data into infra_types for INTERNAL and CUSTOMER
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM infra_types WHERE name = 'INTERNAL') THEN
        INSERT INTO infra_types (name) VALUES ('INTERNAL'), ('CUSTOMER');
    END IF;
END $$;

-- 8. Insert sections for INTERNAL and CUSTOMER infra types
DO $$
BEGIN
    -- Sections for INTERNAL type (Product 1, Product 2)
    IF NOT EXISTS (SELECT 1 FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL')) THEN
        INSERT INTO sections (name, infra_type_id)
        VALUES
            ('Product 1', (SELECT id FROM infra_types WHERE name = 'INTERNAL')),
            ('Product 2', (SELECT id FROM infra_types WHERE name = 'INTERNAL'));
    END IF;

    -- Sections for CUSTOMER type (Vendor A, Vendor B)
    IF NOT EXISTS (SELECT 1 FROM sections WHERE name = 'Vendor A' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'CUSTOMER')) THEN
        INSERT INTO sections (name, infra_type_id)
        VALUES
            ('Vendor A', (SELECT id FROM infra_types WHERE name = 'CUSTOMER')),
            ('Vendor B', (SELECT id FROM infra_types WHERE name = 'CUSTOMER'));
    END IF;
END $$;

-- 9. Insert environment groups for INTERNAL sections (Product 1, Product 2)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM environment_groups WHERE name = 'DEV' AND section_id = (SELECT id FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))) THEN
        INSERT INTO environment_groups (name, section_id)
        VALUES
            ('DEV', (SELECT id FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),
            ('QA', (SELECT id FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),
            ('CONSULT', (SELECT id FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),
            ('PRESALES', (SELECT id FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),

            ('DEV', (SELECT id FROM sections WHERE name = 'Product 2' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),
            ('QA', (SELECT id FROM sections WHERE name = 'Product 2' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),
            ('STAGING', (SELECT id FROM sections WHERE name = 'Product 2' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL')));
    END IF;
END $$;

-- 10. Insert environments within each environment group for Product 1 and Product 2
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM environments WHERE name = 'DEV') THEN
        INSERT INTO environments (name, env_group_id)
        VALUES
            ('DEV', (SELECT id FROM environment_groups WHERE name = 'DEV' AND section_id = (SELECT id FROM sections WHERE name = 'Product 1'))),
            ('QA', (SELECT id FROM environment_groups WHERE name = 'QA' AND section_id = (SELECT id FROM sections WHERE name = 'Product 1'))),
            ('CONSULT', (SELECT id FROM environment_groups WHERE name = 'CONSULT' AND section_id = (SELECT id FROM sections WHERE name = 'Product 1'))),
            ('PRESALES', (SELECT id FROM environment_groups WHERE name = 'PRESALES' AND section_id = (SELECT id FROM sections WHERE name = 'Product 1'))),

            ('DEV', (SELECT id FROM environment_groups WHERE name = 'DEV' AND section_id = (SELECT id FROM sections WHERE name = 'Product 2'))),
            ('QA', (SELECT id FROM environment_groups WHERE name = 'QA' AND section_id = (SELECT id FROM sections WHERE name = 'Product 2'))),
            ('STAGING', (SELECT id FROM environment_groups WHERE name = 'STAGING' AND section_id = (SELECT id FROM sections WHERE name = 'Product 2')));
    END IF;
END $$;

-- 11. Insert customer-specific products for CUSTOMER sections (Vendor A, Vendor B)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM customer_products WHERE product_name = 'Product 1' AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')) THEN
        INSERT INTO customer_products (customer_id, product_name)
        VALUES
            ((SELECT id FROM customers WHERE name = 'Vendor A'), 'Product 1'),
            ((SELECT id FROM customers WHERE name = 'Vendor A'), 'Product 2'),
            ((SELECT id FROM customers WHERE name = 'Vendor B'), 'Product 1'),
            ((SELECT id FROM customers WHERE name = 'Vendor B'), 'Product 2');
    END IF;
END $$;




---------------------------------------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------------------------------------

-- 1. Create the products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- 2. Create the environments table, linked to products
CREATE TABLE IF NOT EXISTS environments (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL
);

-- 3. Create the environment_details table, linked to environments
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

-- Insert dummy data

-- Insert into products
INSERT INTO products (name) VALUES ('Product 1') ON CONFLICT DO NOTHING;

-- Insert into environments
INSERT INTO environments (product_id, name)
VALUES
    ((SELECT id FROM products WHERE name = 'Product 1'), 'DEV')
    ((SELECT id FROM products WHERE name = 'Product 1'), 'QA')
    ((SELECT id FROM products WHERE name = 'Product 1'), 'CONSULT')
    ((SELECT id FROM products WHERE name = 'Product 1'), 'PRESALES')
ON CONFLICT DO NOTHING;

INSERT INTO environments (product_id, name)
VALUES
    ((SELECT id FROM products WHERE name = 'Product 2'), 'DEV')
    ((SELECT id FROM products WHERE name = 'Product 2'), 'QA')
    ((SELECT id FROM products WHERE name = 'Product 2'), 'STAGING')
ON CONFLICT DO NOTHING;

-- Insert into environment_details
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Dev', 'dev.example.com', '2021-08-19 21:30:00', 'Online', 'Samyak', 'develop-20240821.1', '7.2.0876', 'Testing this env so please check'),
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Smoke', 'smoke.example.com', '2021-08-19 21:30:00', 'Online', 'Samyak', 'develop-20240920.3', '7.2.0876', 'Testing this env so please check'),
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), 'Stage', 'stage.example.com', '2021-08-19 21:30:00', 'Online', 'Samyak', 'develop-20240512.1', '7.2.0876', 'Testing this env so please check')
ON CONFLICT DO NOTHING;

-- Product 1

-- Insert QA environments for Product 1
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

############################################################################################################
############################################################################################################

-- Insert into products
INSERT INTO products (name) VALUES ('Product 2') ON CONFLICT DO NOTHING;

-- Insert into environments
INSERT INTO environments (product_id, name)
VALUES
    ((SELECT id FROM products WHERE name = 'Product 2'), 'DEV')
ON CONFLICT DO NOTHING;

-- Product 2

-- Insert into environment_details
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 2')), 'Dev', 'dev.example.com', '2021-08-20 18:45:00', 'Online', 'Samyak', 'develop-20240822.1', '7.2.0876', 'Testing this Product 2 Dev env'),
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 2')), 'Smoke', 'smoke.example.com', '2021-08-21 10:00:00', 'Online', 'Samyak', 'develop-20240921.3', '7.2.0877', 'Testing this Product 2 Smoke env'),
    ((SELECT id FROM environments WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 2')), 'Stage', 'stage.example.com', '2021-08-22 11:00:00', 'Online', 'Samyak', 'develop-20240513.1', '7.2.0878', 'Testing this Product 2 Stage env')
ON CONFLICT DO NOTHING;

-- Insert QA environments for Product 2
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 2')), 'Manual', 'manual.qa.example.com', '2021-08-25 08:00:00', 'Online', 'Isaac', 'qa-manual-20240920.1', '7.2.0884', 'Manual QA environment for Product 2'),
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 2')), 'Auto', 'auto.qa.example.com', '2021-08-25 09:00:00', 'In Progress', 'Jack', 'qa-auto-20240925.2', '7.2.0885', 'Automated QA environment for Product 2'),
    ((SELECT id FROM environments WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 2')), 'Prelaunch', 'prelaunch.qa.example.com', '2021-08-25 10:00:00', 'Offline', 'Karen', 'qa-prelaunch-20241001.3', '7.2.0886', 'Prelaunch QA environment for Product 2')
ON CONFLICT DO NOTHING;

-- Insert STAGING environments for Product 2
INSERT INTO environment_details (environment_id, name, url, last_updated, status, contact, app_version, db_version, comments)
VALUES
    ((SELECT id FROM environments WHERE name = 'STAGING' AND product_id = (SELECT id FROM products WHERE name = 'Product 2')), 'Release', 'release.staging.example.com', '2021-08-26 08:00:00', 'Online', 'Laura', 'staging-release-20241005.1', '7.2.0887', 'Release environment for Product 2 Staging'),
    ((SELECT id FROM environments WHERE name = 'STAGING' AND product_id = (SELECT id FROM products WHERE name = 'Product 2')), 'Launch', 'launch.staging.example.com', '2021-08-26 09:00:00', 'In Progress', 'Mike', 'staging-launch-20241010.2', '7.2.0888', 'Launch environment for Product 2 Staging'),
    ((SELECT id FROM environments WHERE name = 'STAGING' AND product_id = (SELECT id FROM products WHERE name = 'Product 2')), 'Hotfix', 'hotfix.staging.example.com', '2021-08-26 10:00:00', 'Offline', 'Nina', 'staging-hotfix-20241015.3', '7.2.0889', 'Hotfix environment for Product 2 Staging')
ON CONFLICT DO NOTHING;



############################################################################################################
############################################################################################################

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
