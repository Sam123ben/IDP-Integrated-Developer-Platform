// backend/services/structure/structure.go

package structure

import (
	"fmt"
	"os"
)

func EnsureFolderStructure(baseOutputDir string) error {
	// Create base directory
	if err := os.MkdirAll(baseOutputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %v", baseOutputDir, err)
	}

	return nil
}
