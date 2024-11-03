package environments

import (
    "database/sql"
    "log"
    "time"
)

type Environment struct {
    ID          int
    Name        string
    Description string
    UpdatedAt   time.Time
}

// GetLatestData fetches the latest environment data, ordered by `updated_at`
func GetLatestData(db *sql.DB) ([]Environment, error) {
    query := "SELECT id, name, description, updated_at FROM environments ORDER BY updated_at DESC LIMIT 10"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var data []Environment
    for rows.Next() {
        var env Environment
        if err := rows.Scan(&env.ID, &env.Name, &env.Description, &env.UpdatedAt); err != nil {
            return nil, err
        }
        data = append(data, env)
    }

    // Check for errors during row iteration
    if err = rows.Err(); err != nil {
        return nil, err
    }

    log.Printf("Fetched %d environment records", len(data))
    return data, nil
}