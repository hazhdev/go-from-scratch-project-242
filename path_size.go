package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	res, err := getSize(path, all, recursive)
	if err != nil {
		return "", err
	}
	return formatSize(res, human), nil
}

func getSize(path string, all bool, recursive bool) (int64, error) {
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
		if strings.HasPrefix(entry.Name(), ".") && !all {
			continue
		}

		fullPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			if recursive {
				subDirSize, err := getSize(fullPath, all, recursive)
				if err != nil {
					return 0, err
				}
				totalSize += subDirSize
			}

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

func formatSize(bytes int64, human bool) string {
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
