package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

func createLoggerFile(userTempDir string) (*os.File, error) {
	now := time.Now()
	date := fmt.Sprintf("%s.log", now.Format("2006-01-02"))

	if err := os.MkdirAll(path.Join(userTempDir, "kilo-go"), 0o755); err != nil {
		return nil, err
	}

	fileFullPath := path.Join(userTempDir, "kilo-go", date)
	file, err := os.OpenFile(fileFullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err != nil {
		return nil, err
	}

	return file, nil
}
