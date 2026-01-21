package config

import (
	"flag"
	"fmt"
	"strings"
)

type Config struct {
	DebPath              string
	Repeat               int
	VerifyFiles          []string
	VerifyGeneratedFiles []string
	Mocks                []string
	ShowVersion          bool
}

func ParseArgs(args []string) (*Config, error) {
	fs := flag.NewFlagSet("deb-tester", flag.ContinueOnError)

	var (
		repeat          = fs.Int("repeat", 1, "Number of install/remove cycles")
		verifyFiles     = fs.String("verify-files", "", "Semicolon separated list of files to verify inside the deb")
		verifyGenerated = fs.String("verify-generated-files", "", "Semicolon separated list of script-generated files")
		mocks           = fs.String("mock", "", "Semicolon separated list of commands to mock")
		showVersion     = fs.Bool("version", false, "Show version and exit")
	)

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if fs.NArg() == 0 && !*showVersion {
		return nil, fmt.Errorf("missing deb file argument")
	}

	cfg := &Config{
		DebPath:     fs.Arg(0),
		Repeat:      *repeat,
		ShowVersion: *showVersion,
	}

	if *verifyFiles != "" {
		cfg.VerifyFiles = strings.Split(*verifyFiles, ";")
	}
	if *verifyGenerated != "" {
		cfg.VerifyGeneratedFiles = strings.Split(*verifyGenerated, ";")
	}
	if *mocks != "" {
		cfg.Mocks = strings.Split(*mocks, ";")
	}

	return cfg, nil
}
