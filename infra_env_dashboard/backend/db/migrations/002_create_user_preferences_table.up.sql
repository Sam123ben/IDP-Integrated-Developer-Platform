-- Create the new table for storing theme preferences
CREATE TABLE IF NOT EXISTS user_theme_preferences (
    user_id SERIAL PRIMARY KEY,
    theme VARCHAR(20) NOT NULL DEFAULT 'light'
);