-- Create the new table for storing theme preferences
CREATE TABLE user_theme_preferences (
    user_id SERIAL PRIMARY KEY,
    theme VARCHAR(20) NOT NULL DEFAULT 'light'
);