package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	var totalSize int64

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		entryInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}
		totalSize += entryInfo.Size()
	}
	return totalSize, nil
}

func FormatSize(bytes int64, human bool) string {
	if !human || bytes < 1024 {
		return fmt.Sprintf("%dB", bytes)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	value := float64(bytes)
	unitIndex := 0

	for value >= 1024 && unitIndex < len(units)-1 {
		value /= 1024
		unitIndex++
	}

	return fmt.Sprintf("%.1f%s", value, units[unitIndex])
}
