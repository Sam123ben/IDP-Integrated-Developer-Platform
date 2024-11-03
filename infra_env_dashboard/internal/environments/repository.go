package environments

import (
    "database/sql"
    // other imports
)

type Environment struct {
    ID          int
    Name        string
    Description string
    UpdatedAt   string // or time.Time if preferred
}

// Function to fetch the latest data
func GetLatestData() ([]Environment, error) {
    var data []Environment
    rows, err := db.Query("SELECT id, name, description, updated_at FROM environments ORDER BY updated_at DESC LIMIT 10")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var env Environment
        err := rows.Scan(&env.ID, &env.Name, &env.Description, &env.UpdatedAt)
        if err != nil {
            return nil, err
        }
        data = append(data, env)
    }
    return data, nil
}