-- 1. Drop existing data tables if necessary (optional, based on your needs)
-- This can be uncommented if you need to drop tables each time.
-- DROP TABLE IF EXISTS customer_environment_groups;
-- DROP TABLE IF EXISTS product_environment_groups;
-- DROP TABLE IF EXISTS environment_groups;
-- DROP TABLE IF EXISTS customer_products;
-- DROP TABLE IF EXISTS customers;
-- DROP TABLE IF EXISTS applications;
-- DROP TABLE IF EXISTS environments;
-- DROP TABLE IF EXISTS company;

-- 2. Create the company table if it does not exist
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

-- 3. Create environment_groups table if it does not exist (for INTERNAL grouping)
CREATE TABLE IF NOT EXISTS environment_groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    type VARCHAR(20) NOT NULL DEFAULT 'INTERNAL' -- Can be 'INTERNAL' or 'CUSTOMER'
);

-- 4. Create environments table if it does not exist
CREATE TABLE IF NOT EXISTS environments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    last_updated TIMESTAMP NOT NULL,
    status VARCHAR(50) NOT NULL,
    contact VARCHAR(50),
    app_version VARCHAR(20),
    db_version VARCHAR(20),
    comments TEXT,
    status_class VARCHAR(50),
    environment_group_id INT REFERENCES environment_groups(id) ON DELETE SET NULL
);

-- 5. Create applications table if it does not exist
CREATE TABLE IF NOT EXISTS applications (
    id SERIAL PRIMARY KEY,
    environment_id INT REFERENCES environments(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    version VARCHAR(50),
    status VARCHAR(20) -- e.g., 'green', 'orange', 'red'
);

-- 6. Create customers table if it does not exist
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- 7. Create customer_products table if it does not exist
CREATE TABLE IF NOT EXISTS customer_products (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
    product_name VARCHAR(50) NOT NULL
);

-- 8. Create customer_environment_groups table for CUSTOMER-specific environment grouping
CREATE TABLE IF NOT EXISTS customer_environment_groups (
    id SERIAL PRIMARY KEY,
    customer_product_id INT REFERENCES customer_products(id) ON DELETE CASCADE,
    environment_id INT REFERENCES environments(id) ON DELETE CASCADE
);

-- 9. Insert data into environment_groups for INTERNAL groups (if not already present)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM environment_groups WHERE name = 'DEV') THEN
        INSERT INTO environment_groups (name, type) VALUES ('DEV', 'INTERNAL'), ('QA', 'INTERNAL');
    END IF;
END $$;

-- 10. Insert data into environments with environment_group_id for INTERNAL environments
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM environments WHERE name = 'SMOKE') THEN
        INSERT INTO environments (name, last_updated, status, contact, app_version, db_version, comments, status_class, environment_group_id)
        VALUES
            ('SMOKE', '2021-08-19 21:30:00', 'Failed Deployment', 'Taj', '2021.07.27', '7.2.0555', 'Upgrade in progress', 'card-failed', (SELECT id FROM environment_groups WHERE name = 'DEV')),
            ('DEV', '2021-08-19 21:30:00', 'Deployment In Progress', 'Taj', '2021.07.27', '7.2.0555', 'Upgrade in progress', 'card-in-progress', (SELECT id FROM environment_groups WHERE name = 'DEV')),
            ('STAGE', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Running smoothly', 'card-online', (SELECT id FROM environment_groups WHERE name = 'DEV')),
            ('AUTO', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Automated testing in progress', 'card-online', (SELECT id FROM environment_groups WHERE name = 'QA')),
            ('MANUAL', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Manual testing in progress', 'card-online', (SELECT id FROM environment_groups WHERE name = 'QA')),
            ('PRELAUNCH', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Prelaunch preparations', 'card-online', (SELECT id FROM environment_groups WHERE name = 'QA'));
    END IF;
END $$;

-- 11. Insert customers and link them to products
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM customers WHERE name = 'Vendor A') THEN
        INSERT INTO customers (name) VALUES ('Vendor A');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM customers WHERE name = 'Vendor B') THEN
        INSERT INTO customers (name) VALUES ('Vendor B');
    END IF;

    -- Insert into customer_products
    IF NOT EXISTS (SELECT 1 FROM customer_products WHERE customer_id = (SELECT id FROM customers WHERE name = 'Vendor A') AND product_name = 'Product 1') THEN
        INSERT INTO customer_products (customer_id, product_name)
        VALUES
            ((SELECT id FROM customers WHERE name = 'Vendor A'), 'Product 1'),
            ((SELECT id FROM customers WHERE name = 'Vendor A'), 'Product 2'),
            ((SELECT id FROM customers WHERE name = 'Vendor B'), 'Product 1'),
            ((SELECT id FROM customers WHERE name = 'Vendor B'), 'Product 2');
    END IF;
END $$;

-- 12. Insert environments for CUSTOMER section without grouping them under INTERNAL groups
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM environments WHERE name = 'UAT') THEN
        INSERT INTO environments (name, last_updated, status, contact, app_version, db_version, comments, status_class)
        VALUES
            ('UAT', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'User Acceptance Testing in progress', 'card-online'),
            ('TRAIN', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Training environment', 'card-online'),
            ('SUPPORT', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Support environment', 'card-online'),
            ('PROD', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Production environment', 'card-online');
    END IF;
END $$;

-- 13. Link CUSTOMER environments to specific products in customer_environment_groups table
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM customer_environment_groups WHERE customer_product_id = (SELECT id FROM customer_products WHERE product_name = 'Product 1' AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A'))) THEN
        -- Link Product 1 for Vendor A to UAT, TRAIN, SUPPORT, PROD environments
        INSERT INTO customer_environment_groups (customer_product_id, environment_id)
        VALUES
            ((SELECT id FROM customer_products WHERE product_name = 'Product 1' AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')), (SELECT id FROM environments WHERE name = 'UAT')),
            ((SELECT id FROM customer_products WHERE product_name = 'Product 1' AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')), (SELECT id FROM environments WHERE name = 'TRAIN')),
            ((SELECT id FROM customer_products WHERE product_name = 'Product 1' AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')), (SELECT id FROM environments WHERE name = 'SUPPORT')),
            ((SELECT id FROM customer_products WHERE product_name = 'Product 1' AND customer_id = (SELECT id FROM customers WHERE name = 'Vendor A')), (SELECT id FROM environments WHERE name = 'PROD'));

        -- Repeat similar entries for Vendor B and other products if needed
    END IF;
END $$;

-- 14. Insert default applications into each environment if not already present
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM applications WHERE environment_id = (SELECT id FROM environments WHERE name = 'SMOKE') AND name = 'portal') THEN
        -- Applications for SMOKE environment
        INSERT INTO applications (environment_id, name, version, status)
        VALUES
            ((SELECT id FROM environments WHERE name = 'SMOKE'), 'portal', 'develop-20240201', 'green'),
            ((SELECT id FROM environments WHERE name = 'SMOKE'), 'idsrv', 'develop-20231113', 'orange'),
            ((SELECT id FROM environments WHERE name = 'SMOKE'), 'bis', 'develop-20240120', 'red'),
            ((SELECT id FROM environments WHERE name = 'SMOKE'), 'analytics', 'develop-20231215', 'green');

        -- Applications for DEV environment
        INSERT INTO applications (environment_id, name, version, status)
        VALUES
            ((SELECT id FROM environments WHERE name = 'DEV'), 'portal', 'develop-20240201', 'green'),
            ((SELECT id FROM environments WHERE name = 'DEV'), 'idsrv', 'develop-20231113', 'orange'),
            ((SELECT id FROM environments WHERE name = 'DEV'), 'bis', 'develop-20240120', 'green'),
            ((SELECT id FROM environments WHERE name = 'DEV'), 'analytics', 'develop-20231215', 'red');

        -- Applications for STAGE environment
        INSERT INTO applications (environment_id, name, version, status)
        VALUES
            ((SELECT id FROM environments WHERE name = 'STAGE'), 'portal', 'develop-20240201', 'green'),
            ((SELECT id FROM environments WHERE name = 'STAGE'), 'idsrv', 'develop-20231113', 'green'),
            ((SELECT id FROM environments WHERE name = 'STAGE'), 'bis', 'develop-20240120', 'green'),
            ((SELECT id FROM environments WHERE name = 'STAGE'), 'analytics', 'develop-20231215', 'orange');

        -- Repeat similar application inserts for CUSTOMER environments like UAT, TRAIN, SUPPORT, PROD
    END IF;
END $$;