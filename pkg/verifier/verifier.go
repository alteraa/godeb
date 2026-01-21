package verifier

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"deb-tester/pkg/logger"
)

func VerifyDebContent(debPath string) error {
	logger.Info("Verifying contents of %s...", debPath)
	// dpkg -c lists contents
	cmd := exec.Command("dpkg", "-c", debPath)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("dpkg -c failed: %v", err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	verifiedCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		
		parts := strings.Fields(line)
		if len(parts) < 6 {
			continue
		}
		path := parts[len(parts)-1]
		
		// Remove leading dot
		absPath := strings.TrimPrefix(path, ".")
		
		// Skip directories (ending in /)
		if strings.HasSuffix(absPath, "/") {
			continue
		}

		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			return fmt.Errorf("expected file missing: %s", absPath)
		}
		logger.Debug("Verified file exists: %s", absPath)
		verifiedCount++
	}
	logger.Success("Successfully verified %d files from %s", verifiedCount, debPath)

	return nil
}

func VerifyGeneratedFiles(files []string) error {
	if len(files) == 0 {
		return nil
	}
	logger.Info("Verifying %d generated files...", len(files))
	for _, f := range files {
		if _, err := os.Stat(f); os.IsNotExist(err) {
			return fmt.Errorf("script-generated file missing: %s", f)
		}
		logger.Debug("Verified generated file exists: %s", f)
	}
	return nil
}

func VerifyClean(debPath string, generatedFiles []string) error {
	logger.Info("Verifying cleanup for %s...", debPath)
	// 1. Check deb contents are gone
	cmd := exec.Command("dpkg", "-c", debPath)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("dpkg -c failed during cleanup check: %v", err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	cleanedCount := 0
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) < 6 { continue }
		path := parts[len(parts)-1]
		absPath := strings.TrimPrefix(path, ".")
		
		if strings.HasSuffix(absPath, "/") { continue }

		if _, err := os.Stat(absPath); !os.IsNotExist(err) {
			return fmt.Errorf("file remains after removal: %s", absPath)
		}
		logger.Debug("Verified file removed: %s", absPath)
		cleanedCount++
	}
	logger.Success("Verified %d package files were cleaned up.", cleanedCount)

	// 2. Check generated files are gone
	if len(generatedFiles) > 0 {
		logger.Info("Verifying cleanup of %d generated files...", len(generatedFiles))
		for _, f := range generatedFiles {
			if _, err := os.Stat(f); !os.IsNotExist(err) {
				return fmt.Errorf("generated file remains after removal: %s", f)
			}
			logger.Debug("Verified generated file removed: %s", f)
		}
	}

	return nil
}

