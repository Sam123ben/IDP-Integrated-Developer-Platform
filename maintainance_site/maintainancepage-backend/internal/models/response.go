// backend/internal/models/response.go
package models

type MaintenanceResponse struct {
	CurrentTimeUTC   string                `json:"current_time_utc"`
	CurrentTimeLocal string                `json:"current_time_local"`
	IsActive         bool                  `json:"is_active"`
	RemainingTime    int                   `json:"remaining_time_minutes"`
	Maintenance      MaintenanceWindowData `json:"maintenance"`
}

type MaintenanceWindowData struct {
	ID          uint                         `json:"id"`
	StartTime   string                       `json:"start_time"`
	Description string                       `json:"description"`
	Components  map[string][]SystemComponent `json:"components"`
	Updates     PaginatedUpdates             `json:"updates"`
}

type PaginatedUpdates struct {
	Data       []MaintenanceUpdate `json:"data"`
	Pagination Pagination          `json:"pagination"`
}

type Pagination struct {
	CurrentPage  int `json:"current_page"`
	PerPage      int `json:"per_page"`
	TotalUpdates int `json:"total_updates"`
}
