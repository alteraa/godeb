package runner

import (
	"fmt"

	"deb-tester/pkg/config"
	"deb-tester/pkg/installer"
	"deb-tester/pkg/logger"
	"deb-tester/pkg/mocker"
	"deb-tester/pkg/verifier"
)

func Run(cfg *config.Config) error {
	// 1. Setup Mocks
	mockManager := mocker.New(cfg.Mocks)
	if err := mockManager.Setup(); err != nil {
		return fmt.Errorf("failed to setup mocks: %w", err)
	}
	defer mockManager.Cleanup()

	// 2. Loop
	for i := 1; i <= cfg.Repeat; i++ {
		logger.Header("Iteration %d of %d", i, cfg.Repeat)

		// Install
		logger.Phase("[Phase 1/4] Installing %s", cfg.DebPath)
		if err := installer.Install(cfg.DebPath); err != nil {
			return fmt.Errorf("install failed: %w", err)
		}

		// Verify
		logger.Phase("[Phase 2/4] Verifying Installation")
		if err := verifier.VerifyDebContent(cfg.DebPath); err != nil {
			return fmt.Errorf("deb content verification failed: %w", err)
		}
		if err := verifier.VerifyGeneratedFiles(cfg.VerifyGeneratedFiles); err != nil {
			return fmt.Errorf("generated file verification failed: %w", err)
		}

		// Remove
		logger.Phase("[Phase 3/4] Removing Package")
		if err := installer.Remove(cfg.DebPath); err != nil {
			return fmt.Errorf("remove failed: %w", err)
		}

		// Verify Clean
		logger.Phase("[Phase 4/4] Verifying Cleanup")
		if err := verifier.VerifyClean(cfg.DebPath, cfg.VerifyGeneratedFiles); err != nil {
			return fmt.Errorf("cleanup verification failed: %w", err)
		}

		if err := mockManager.VerifyUsage(); err != nil {
			return fmt.Errorf("mock usage verification failed: %w", err)
		}

		logger.Success("Iteration %d completed successfully.", i)
	}

	return nil
}
