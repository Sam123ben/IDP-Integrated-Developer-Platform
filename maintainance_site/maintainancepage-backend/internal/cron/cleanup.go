// backend/cron/cleanup.go

package croncleanup

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// CleanupMaintenanceUpdates deletes updates for fixed maintenance windows
func CleanupMaintenanceUpdates(db *gorm.DB) error {
	log.Println("Running cleanup for fixed maintenance windows")
	result := db.Exec(`
		DELETE FROM maintenance_updates
		WHERE maintenance_window_id IN (
			SELECT id
			FROM maintenance_windows
			WHERE status = 'fixed' AND issue_fixed = true
		)
	`)
	if result.Error != nil {
		log.Printf("Error cleaning up maintenance updates: %v", result.Error)
		return result.Error
	}

	log.Printf("Deleted %d maintenance updates for fixed windows", result.RowsAffected)
	return nil
}

// StartCleanupCron schedules the cleanup function to run every 30 minutes
func StartCleanupCron(db *gorm.DB) {
	ticker := time.NewTicker(30 * time.Minute)

	go func() {
		for range ticker.C {
			log.Println("Cron job triggered for maintenance updates cleanup")
			if err := CleanupMaintenanceUpdates(db); err != nil {
				log.Printf("Error during maintenance updates cleanup: %v", err)
			}
		}
	}()
}
