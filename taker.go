package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type NoteTaker struct {
	baseDir    string
	dateFormat string
	fileMode   os.FileMode
}

func NewNoteTaker(baseDir string) *NoteTaker {
	return &NoteTaker{
		baseDir:    baseDir,
		dateFormat: "2006-01-02",
		fileMode:   0644,
	}
}

func (nt *NoteTaker) formatNote(content string) string {
	return fmt.Sprintf("\n- [ ] %s", strings.TrimSpace(content))
}

func (nt *NoteTaker) getTodayFilePath() string {
	fileName := fmt.Sprintf("%s.md", time.Now().Format(nt.dateFormat))
	return filepath.Join(nt.baseDir, fileName)
}

func (nt *NoteTaker) WriteNote(content string) error {
	if content == "" {
		return fmt.Errorf("note content cannot be empty")
	}

	if err := os.MkdirAll(nt.baseDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePath := nt.getTodayFilePath()
	formattedNote := nt.formatNote(content)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, nt.fileMode)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	if _, err := file.WriteString(formattedNote); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
