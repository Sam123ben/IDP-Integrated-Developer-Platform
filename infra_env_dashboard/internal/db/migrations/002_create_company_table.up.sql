-- Create the company table if it does not exist
CREATE TABLE IF NOT EXISTS company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Delete any existing records from the company table
DELETE FROM company;

-- Insert the new company record
INSERT INTO company (name)
VALUES ('My Company');