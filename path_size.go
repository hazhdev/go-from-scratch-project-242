package code

import (
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