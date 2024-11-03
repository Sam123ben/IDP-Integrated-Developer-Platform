package environments

import (
    "database/sql"
    "infra_env_dashboard/internal/models"
    "log"
)

// GetLatestData fetches the latest environment data with the new columns
func GetLatestData(db *sql.DB) ([]models.Environment, error) {
    query := `
    SELECT id, environment_name, description, created_at, updated_at, environment_type,
           group_name, customer_name, url, status, contact, app_version, db_version, comments
    FROM environments
    ORDER BY updated_at DESC LIMIT 10`
    
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var data []models.Environment
    for rows.Next() {
        var env models.Environment
        if err := rows.Scan(&env.ID, &env.EnvironmentName, &env.Description, &env.CreatedAt,
                            &env.UpdatedAt, &env.EnvironmentType, &env.GroupName,
                            &env.CustomerName, &env.URL, &env.Status, &env.Contact,
                            &env.AppVersion, &env.DBVersion, &env.Comments); err != nil {
            return nil, err
        }
        data = append(data, env)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    log.Printf("Fetched %d environment records", len(data))
    return data, nil
}