package mocker

import (
	"fmt"
	"os"
	"path/filepath"

	"deb-tester/pkg/logger"
)

const (
	MockDir = "/usr/bin"
	MockLog = "/tmp/mock_log.txt"
)

type Manager struct {
	mocks []string
}

func New(mocks []string) *Manager {
	return &Manager{mocks: mocks}
}

func (m *Manager) Setup() error {
	if len(m.mocks) == 0 {
		return nil
	}

	logger.Info("Setting up %d mocks in %s", len(m.mocks), MockDir)

	// Create log file
	f, err := os.Create(MockLog)
	if err != nil {
		return fmt.Errorf("failed to create mock log: %v", err)
	}
	f.Close()
	os.Chmod(MockLog, 0666) // Allow writing by anyone
	logger.Debug("Initialized mock log at %s", MockLog)

	// Create mock dir
	if err := os.MkdirAll(MockDir, 0755); err != nil {
		return fmt.Errorf("failed to create mock dir: %v", err)
	}

	// Create shim scripts
	for _, mock := range m.mocks {
		path := filepath.Join(MockDir, mock)
		content := fmt.Sprintf("#!/bin/sh\necho \"%s $@\" >> %s\n", mock, MockLog)
		
		if err := os.WriteFile(path, []byte(content), 0755); err != nil {
			return fmt.Errorf("failed to create mock %s: %v", mock, err)
		}
		logger.Debug("Created mock shim: %s", path)
	}

	// Prepend to PATH
	currentPath := os.Getenv("PATH")
	newPath := MockDir + ":" + currentPath
	os.Setenv("PATH", newPath)
	logger.Debug("Prepended %s to PATH", MockDir)

	return nil
}

func (m *Manager) VerifyUsage() error {
	if len(m.mocks) == 0 {
		return nil
	}
	
	content, err := os.ReadFile(MockLog)
	if err != nil {
		return fmt.Errorf("failed to read mock log: %v", err)
	}
	
	logger.Info("--- Mock Usage Log Start ---\n%s--- Mock Usage Log End ---", string(content))
	return nil
}

func (m *Manager) Cleanup() {
	if len(m.mocks) == 0 {
		return
	}
	logger.Info("Cleaning up mock shims in %s", MockDir)
	for _, mock := range m.mocks {
		path := filepath.Join(MockDir, mock)
		os.Remove(path)
	}
	os.Remove(MockLog)
}

