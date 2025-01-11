// backend/internal/models/models.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type SystemStatus string
type MaintenanceStatus string

const (
	StatusOperational SystemStatus      = "operational"
	StatusMaintenance SystemStatus      = "maintenance"
	StatusDegraded    SystemStatus      = "degraded"
	StatusActive      MaintenanceStatus = "active"
	StatusFixed       MaintenanceStatus = "fixed"
)

type AuditEntry struct {
	gorm.Model
	MaintenanceWindowID uint      `json:"maintenance_window_id" gorm:"index;not null"`
	Action              string    `json:"action" gorm:"not null"`
	PerformedBy         string    `json:"performed_by" gorm:"not null"`
	Timestamp           time.Time `json:"timestamp" gorm:"type:timestamp with time zone;not null"`
}

type MaintenanceWindow struct {
	gorm.Model
	StartTime         time.Time           `json:"start_time" gorm:"type:timestamp with time zone;index;not null"`
	EstimatedDuration int                 `json:"estimated_duration" gorm:"type:bigint;not null"`
	Description       string              `json:"description"`
	Status            MaintenanceStatus   `json:"status" gorm:"type:maintenance_status;default:'active'"`
	IssueFixed        bool                `json:"issue_fixed" gorm:"default:false"`
	Components        []SystemComponent   `json:"components" gorm:"many2many:maintenance_components;constraint:OnDelete:CASCADE;"`
	Updates           []MaintenanceUpdate `json:"updates" gorm:"foreignKey:MaintenanceWindowID;constraint:OnDelete:CASCADE;"`
}

// BeforeCreate ensures time is stored in UTC
func (m *MaintenanceWindow) BeforeCreate(tx *gorm.DB) error {
	m.StartTime = m.StartTime.UTC()
	return nil
}

// BeforeUpdate ensures time is stored in UTC during updates
func (m *MaintenanceWindow) BeforeUpdate(tx *gorm.DB) error {
	m.StartTime = m.StartTime.UTC()
	return nil
}

// AfterFind keeps the time in UTC
func (m *MaintenanceWindow) AfterFind(tx *gorm.DB) error {
	m.StartTime = m.StartTime.UTC()
	return nil
}

type SystemComponent struct {
	gorm.Model
	Name   string       `json:"name" gorm:"uniqueIndex;not null"`
	Type   string       `json:"type"`
	Status SystemStatus `json:"status" gorm:"type:system_status;default:'operational'"`
}

type MaintenanceUpdate struct {
	gorm.Model
	MaintenanceWindowID uint   `json:"-" gorm:"index"`
	Message             string `json:"message"`
}

// AutoMigrate performs automatic database schema migration for all models
func AutoMigrate(db *gorm.DB) error {
	db.Exec(`SET TIME ZONE 'UTC';`)

	// Create enums if they don't exist
	db.Exec(`
    DO $$ 
    BEGIN
        CREATE TYPE system_status AS ENUM ('operational', 'maintenance', 'degraded');
        CREATE TYPE maintenance_status AS ENUM ('active', 'fixed');
    EXCEPTION
        WHEN duplicate_object THEN NULL;
    END $$;
	`)

	// Add issue_fixed column if it doesn't exist
	db.Exec(`
    DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1 FROM information_schema.columns 
            WHERE table_name = 'maintenance_windows' AND column_name = 'issue_fixed'
        ) THEN
            ALTER TABLE maintenance_windows ADD COLUMN issue_fixed BOOLEAN DEFAULT false;
        END IF;
    END $$;
	`)

	return db.AutoMigrate(&MaintenanceWindow{}, &SystemComponent{}, &MaintenanceUpdate{})
}
func (m *MaintenanceWindow) EndTime() time.Time {
	return m.StartTime.UTC().Add(time.Duration(m.EstimatedDuration) * time.Minute)
}

func (m *MaintenanceWindow) IsActive() bool {
	return m.Status == StatusActive
}

func (m *MaintenanceWindow) RemainingMinutes(currentTime time.Time) int {
	endTime := m.EndTime()
	remaining := endTime.Sub(currentTime).Minutes()
	if remaining < 0 {
		return 0
	}
	return int(remaining)
}
