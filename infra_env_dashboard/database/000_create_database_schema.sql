-- 1. Drop existing data tables if necessary (optional)
-- DROP TABLE IF EXISTS customer_environment_groups;
-- DROP TABLE IF EXISTS product_environment_groups;
-- DROP TABLE IF EXISTS infra_types;
-- DROP TABLE IF EXISTS environment_groups;
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

-- 5. Create env_groups table to hold environment group names (e.g., DEV, QA, CONSULT)
CREATE TABLE IF NOT EXISTS env_groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE, -- Environment group name
    product_id INT REFERENCES products(id) ON DELETE CASCADE -- Linked to specific product under infra type
);

-- 6. Create environments table for detailed environments within each group
CREATE TABLE IF NOT EXISTS environments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL, -- Environment name, e.g., SMOKE, STAGE, etc.
    last_updated TIMESTAMP NOT NULL,
    status VARCHAR(50) NOT NULL,
    contact VARCHAR(50),
    app_version VARCHAR(20),
    db_version VARCHAR(20),
    comments TEXT,
    status_class VARCHAR(50),
    env_group_id INT REFERENCES env_groups(id) ON DELETE SET NULL -- Linked to an env_group
);

-- 7. Create applications table for applications within each environment
CREATE TABLE IF NOT EXISTS applications (
    id SERIAL PRIMARY KEY,
    environment_id INT REFERENCES environments(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    version VARCHAR(50),
    status VARCHAR(20) -- e.g., 'green', 'orange', 'red'
);

-- 8. Create customers table to hold customer names
CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- 9. Create customer_products table to link customers to products
CREATE TABLE IF NOT EXISTS customer_products (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id) ON DELETE CASCADE
);

-- 10. Create customer_environment_groups table for CUSTOMER environments per product and customer
CREATE TABLE IF NOT EXISTS customer_environment_groups (
    id SERIAL PRIMARY KEY,
    customer_product_id INT REFERENCES customer_products(id) ON DELETE CASCADE,
    environment_id INT REFERENCES environments(id) ON DELETE CASCADE
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

-- 13. Insert env_groups for INTERNAL products (DEV, QA, CONSULT)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM env_groups WHERE name = 'DEV') THEN
        INSERT INTO env_groups (name, product_id)
        VALUES
            ('DEV', (SELECT id FROM products WHERE name = 'Product 1')),
            ('QA', (SELECT id FROM products WHERE name = 'Product 1')),
            ('CONSULT', (SELECT id FROM products WHERE name = 'Product 1')),
            ('DEV', (SELECT id FROM products WHERE name = 'Product 2')),
            ('QA', (SELECT id FROM products WHERE name = 'Product 2')),
            ('CONSULT', (SELECT id FROM products WHERE name = 'Product 2'));
    END IF;
END $$;

-- 14. Insert environments under each env_group
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM environments WHERE name = 'SMOKE') THEN
        INSERT INTO environments (name, last_updated, status, contact, app_version, db_version, comments, status_class, env_group_id)
        VALUES
            -- DEV environments for Product 1
            ('SMOKE', '2021-08-19 21:30:00', 'Failed Deployment', 'Taj', '2021.07.27', '7.2.0555', 'Upgrade in progress', 'card-failed', (SELECT id FROM env_groups WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))),
            ('DEV', '2021-08-19 21:30:00', 'Deployment In Progress', 'Taj', '2021.07.27', '7.2.0555', 'Upgrade in progress', 'card-in-progress', (SELECT id FROM env_groups WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))),
            ('STAGE', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Running smoothly', 'card-online', (SELECT id FROM env_groups WHERE name = 'DEV' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))),
            
            -- QA environments for Product 1
            ('AUTO', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Automated testing in progress', 'card-online', (SELECT id FROM env_groups WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))),
            ('MANUAL', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Manual testing in progress', 'card-online', (SELECT id FROM env_groups WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))),
            ('PRELAUNCH', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Prelaunch preparations', 'card-online', (SELECT id FROM env_groups WHERE name = 'QA' AND product_id = (SELECT id FROM products WHERE name = 'Product 1')));
    END IF;
END $$;

-- 15. Insert customers and link them to products for CUSTOMER type
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM customers WHERE name = 'Vendor A') THEN
        INSERT INTO customers (name) VALUES ('Vendor A');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM customers WHERE name = 'Vendor B') THEN
        INSERT INTO customers (name) VALUES ('Vendor B');
    END IF;

    -- Link each customer to products under CUSTOMER infra type
    IF NOT EXISTS (SELECT 1 FROM customer_products WHERE customer_id = (SELECT id FROM customers WHERE name = 'Vendor A') AND product_id = (SELECT id FROM products WHERE name = 'Product 1')) THEN
        INSERT INTO customer_products (customer_id, product_id)
        VALUES
            ((SELECT id FROM customers WHERE name = 'Vendor A'), (SELECT id FROM products WHERE name = 'Product 1')),
            ((SELECT id FROM customers WHERE name = 'Vendor A'), (SELECT id FROM products WHERE name = 'Product 2')),
            ((SELECT id FROM customers WHERE name = 'Vendor B'), (SELECT id FROM products WHERE name = 'Product 1')),
            ((SELECT id FROM customers WHERE name = 'Vendor B'), (SELECT id FROM products WHERE name = 'Product 2'));
    END IF;
END $$;

-- 16. Insert environments for CUSTOMER type (e.g., UAT, TRAIN, SUPPORT, PROD)
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM environments WHERE name = 'UAT Environment') THEN
        INSERT INTO environments (name, last_updated, status, contact, app_version, db_version, comments, status_class)
        VALUES
            ('UAT Environment', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'User Acceptance Testing in progress', 'card-online'),
            ('TRAIN Environment', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Training environment', 'card-online'),
            ('SUPPORT Environment', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Support environment', 'card-online'),
            ('PROD Environment', '2021-08-19 21:30:00', 'Online', 'Taj', '2021.07.27', '7.2.0555', 'Production environment', 'card-online');
    END IF;
END $$;

-- 17. Link CUSTOMER environments to specific products and customers in customer_environment_groups
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM customer_environment_groups WHERE customer_product_id = (SELECT id FROM customer_products WHERE customer_id = (SELECT id FROM customers WHERE name = 'Vendor A') AND product_id = (SELECT id FROM products WHERE name = 'Product 1'))) THEN
        INSERT INTO customer_environment_groups (customer_product_id, environment_id)
        VALUES
            ((SELECT id FROM customer_products WHERE customer_id = (SELECT id FROM customers WHERE name = 'Vendor A') AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), (SELECT id FROM environments WHERE name = 'UAT Environment')),
            ((SELECT id FROM customer_products WHERE customer_id = (SELECT id FROM customers WHERE name = 'Vendor A') AND product_id = (SELECT id FROM products WHERE name = 'Product 1')), (SELECT id FROM environments WHERE name = 'PROD Environment'));
    END IF;
END $$;