package filescanner

import (
	"fmt"
	"os"
	"path/filepath"
)

func ScanDir(dirPath string, fileChan chan<- string) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("No se pudo leer el directorio %s", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			fullPath := filepath.Join(dirPath, entry.Name())
			fileChan <- fullPath
		}
	}
	return nil
}