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

-- 2. Create the company table
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

-- 3. Create infra_types table to hold INTERNAL and CUSTOMER types
CREATE TABLE IF NOT EXISTS infra_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE -- e.g., INTERNAL, CUSTOMER
);

-- 4. Create products table, linked to infra_type
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE, -- Product name, e.g., Product 1, Product 2
    infra_type_id INT REFERENCES infra_types(id) ON DELETE CASCADE
);

-- 5. Create sections table to store sections under each infrastructure type
-- Sections will hold either products for INTERNAL or customers for CUSTOMER
CREATE TABLE IF NOT EXISTS sections (
    id SERIAL PRIMARY KEY,
    infra_type_id INTEGER REFERENCES infra_types(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL, -- e.g., Product 1, Vendor A, Vendor B
    product_id INTEGER REFERENCES products(id) ON DELETE SET NULL, -- Product reference for INTERNAL infra type
    customer_id INTEGER REFERENCES customers(id) ON DELETE SET NULL -- Customer reference for CUSTOMER infra type
);

-- 6. Create env_groups table to hold environment group names (e.g., DEV, QA, CONSULT)
-- Each environment group will be linked to a section within INTERNAL or CUSTOMER
CREATE TABLE IF NOT EXISTS env_groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL, -- Environment group name (e.g., DEV, QA, CONSULT, UAT, PROD)
    section_id INT REFERENCES sections(id) ON DELETE CASCADE -- Linked to specific section (product or customer section)
);

-- 7. Create environments table for detailed environments within each group
CREATE TABLE IF NOT EXISTS environments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL, -- Environment name, e.g., SMOKE, STAGE, UAT, PROD, etc.
    last_updated TIMESTAMP NOT NULL,
    status VARCHAR(50) NOT NULL,
    contact VARCHAR(50),
    app_version VARCHAR(20),
    db_version VARCHAR(20),
    comments TEXT,
    status_class VARCHAR(50),
    env_group_id INT REFERENCES env_groups(id) ON DELETE SET NULL -- Linked to an env_group
);

-- 8. Create applications table for applications within each environment
CREATE TABLE IF NOT EXISTS applications (
    id SERIAL PRIMARY KEY,
    environment_id INT REFERENCES environments(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    version VARCHAR(50),
    status VARCHAR(20) -- e.g., 'green', 'orange', 'red'
);

-- 9. Create customers table to hold customer names
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- 10. Create customer_products table to link customers to products under CUSTOMER infra type
CREATE TABLE IF NOT EXISTS customer_products (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id) ON DELETE CASCADE
);

-- 11. Insert data into infra_types for INTERNAL and CUSTOMER types
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM infra_types WHERE name = 'INTERNAL') THEN
        INSERT INTO infra_types (name) VALUES ('INTERNAL'), ('CUSTOMER');
    END IF;
END $$;

-- 12. Insert products for INTERNAL type
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM products WHERE name = 'Product 1') THEN
        INSERT INTO products (name, infra_type_id)
        VALUES
            ('Product 1', (SELECT id FROM infra_types WHERE name = 'INTERNAL')),
            ('Product 2', (SELECT id FROM infra_types WHERE name = 'INTERNAL'));
    END IF;
END $$;

-- 13. Insert sections for INTERNAL products and CUSTOMER vendors
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL')) THEN
        INSERT INTO sections (name, infra_type_id, product_id)
        VALUES
            ('Product 1', (SELECT id FROM infra_types WHERE name = 'INTERNAL'), (SELECT id FROM products WHERE name = 'Product 1')),
            ('Product 2', (SELECT id FROM infra_types WHERE name = 'INTERNAL'), (SELECT id FROM products WHERE name = 'Product 2'));
    END IF;

    IF NOT EXISTS (SELECT 1 FROM sections WHERE name = 'Vendor A' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'CUSTOMER')) THEN
        INSERT INTO sections (name, infra_type_id, customer_id)
        VALUES
            ('Vendor A', (SELECT id FROM infra_types WHERE name = 'CUSTOMER'), (SELECT id FROM customers WHERE name = 'Vendor A')),
            ('Vendor B', (SELECT id FROM infra_types WHERE name = 'CUSTOMER'), (SELECT id FROM customers WHERE name = 'Vendor B'));
    END IF;
END $$;

-- 14. Insert env_groups for INTERNAL product sections (e.g., DEV, QA, CONSULT)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM env_groups WHERE name = 'DEV') THEN
        INSERT INTO env_groups (name, section_id)
        VALUES
            ('DEV', (SELECT id FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),
            ('QA', (SELECT id FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),
            ('CONSULT', (SELECT id FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),
            ('DEV', (SELECT id FROM sections WHERE name = 'Product 2' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),
            ('QA', (SELECT id FROM sections WHERE name = 'Product 2' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))),
            ('CONSULT', (SELECT id FROM sections WHERE name = 'Product 2' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL')));
    END IF;
END $$;

-- 15. Insert environments under each env_group for INTERNAL type
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM environments WHERE name = 'SMOKE') THEN
        INSERT INTO environments (name, last_updated, status, contact, app_version, db_version, comments, status_class, env_group_id)
        VALUES
            ('SMOKE', '2021-08-19 21:30:00', 'Failed Deployment', 'Taj', '2021.07.27', '7.2.0555', 'Upgrade in progress', 'card-failed', (SELECT id FROM env_groups WHERE name = 'DEV' AND section_id = (SELECT id FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL')))),
            ('STAGE', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Running smoothly', 'card-online', (SELECT id FROM env_groups WHERE name = 'DEV' AND section_id = (SELECT id FROM sections WHERE name = 'Product 1' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'INTERNAL'))));
    END IF;
END $$;

-- 16. Insert environments for CUSTOMER type (e.g., UAT, PROD) within Vendor sections
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM environments WHERE name = 'UAT Environment') THEN
        INSERT INTO environments (name, last_updated, status, contact, app_version, db_version, comments, status_class, env_group_id)
        VALUES
            ('UAT Environment', '2021-08-19 21:30:00', 'User Acceptance Testing', 'Taj', '2021.07.27', '7.2.0555', 'UAT phase ongoing', 'card-online', (SELECT id FROM env_groups WHERE name = 'UAT' AND section_id = (SELECT id FROM sections WHERE name = 'Vendor A' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'CUSTOMER')))),
            ('PROD Environment', '2021-08-19 21:30:00', 'Production', 'Taj', '2021.07.27', '7.2.0555', 'Live environment', 'card-online', (SELECT id FROM env_groups WHERE name = 'PROD' AND section_id = (SELECT id FROM sections WHERE name = 'Vendor A' AND infra_type_id = (SELECT id FROM infra_types WHERE name = 'CUSTOMER'))));
    END IF;
END $$;

-- 17. Insert applications into environments
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM applications WHERE name = 'Portal' AND environment_id = (SELECT id FROM environments WHERE name = 'SMOKE')) THEN
        INSERT INTO applications (environment_id, name, version, status)
        VALUES
            ((SELECT id FROM environments WHERE name = 'SMOKE'), 'Portal', 'develop-20240201', 'green'),
            ((SELECT id FROM environments WHERE name = 'STAGE'), 'Portal', 'develop-20240201', 'green');
    END IF;
END $$;