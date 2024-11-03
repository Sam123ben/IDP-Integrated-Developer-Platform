-- This file is used to rollback the changes made in the up.sql file
-- 000_create_company_table.down.sql

-- Drop the company table if it exists
DROP TABLE IF EXISTS company;

-- Optionally, drop the customers table if it was previously used
DROP TABLE IF EXISTS customers;

