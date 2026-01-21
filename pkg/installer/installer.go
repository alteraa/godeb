package installer

import (
	"fmt"
	"os/exec"
	"strings"

	"deb-tester/pkg/logger"
)

func Install(debPath string) error {
	// apt-get install ./package.deb
	// We use ./ prefix to ensure apt treats it as a file
	if !strings.HasPrefix(debPath, "/") && !strings.HasPrefix(debPath, "./") {
		debPath = "./" + debPath
	}

	logger.Info("Executing: apt-get install -y %s", debPath)
	cmd := exec.Command("apt-get", "install", "-y", debPath)
	out, err := cmd.CombinedOutput()
	if out != nil {
		logger.Debug("Output: %s", string(out))
	}
	if err != nil {
		return fmt.Errorf("apt-get install failed: %v, output: %s", err, out)
	}
	logger.Success("Installation successful for %s", debPath)
	return nil
}

func Remove(debPath string) error {
	// 1. Get package name from deb file
	// dpkg-deb -f package.deb Package
	logger.Info("Extracting package name from %s", debPath)
	cmdName := exec.Command("dpkg-deb", "-f", debPath, "Package")
	nameOut, err := cmdName.Output()
	if err != nil {
		return fmt.Errorf("failed to get package name: %v", err)
	}
	pkgName := strings.TrimSpace(string(nameOut))

	if pkgName == "" {
		return fmt.Errorf("could not determine package name from %s", debPath)
	}
	logger.Info("Package name identified: %s", pkgName)

	// 2. Remove
	logger.Info("Executing: apt-get remove -y --purge %s", pkgName)
	cmdRemove := exec.Command("apt-get", "remove", "-y", "--purge", pkgName)
	out, err := cmdRemove.CombinedOutput()
	if err != nil {
		return fmt.Errorf("apt-get remove failed: %v, output: %s", err, out)
	}
	logger.Success("Removal successful for %s", pkgName)
	return nil
}

