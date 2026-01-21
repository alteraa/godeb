package main

import (
	"errors"
	"flag"
	"os"

	"deb-tester/pkg/config"
	"deb-tester/pkg/logger"
	"deb-tester/pkg/runner"
)

func main() {
	cfg, err := config.ParseArgs(os.Args[1:])
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			os.Exit(0)
		}
		logger.Error("Configuration error: %v", err)
		os.Exit(1)
	}

	logger.Info("Starting Deb Tester with: %+v", cfg)

	if err := runner.Run(cfg); err != nil {
		logger.Error("Test failed: %v", err)
		os.Exit(1)
	}

	logger.Success("Test completed successfully.")
}
