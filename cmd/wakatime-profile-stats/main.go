package main

import (
	"fmt"
	"log"
	"os"

	// This controls the maxprocs environment variable in container runtimes.
	// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/zap"
)

func main() {
	env := os.Getenv("ENV")
	var logger *zap.Logger
	var err error
	if env == "prod" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Fatalf("failed to setup logger: %v", err)
	}
	zap.ReplaceGlobals(logger)

	if err := run(); err != nil {
		zap.L().Error("an error occurred", zap.Error(err))
		os.Exit(1)
	}
}

func run() error {
	_, err := maxprocs.Set(maxprocs.Logger(func(s string, i ...interface{}) {
		zap.L().Debug(fmt.Sprintf(s, i...))
	}))
	if err != nil {
		return fmt.Errorf("setting max procs: %w", err)
	}

	zap.L().Info("Starting Wakatime Profile Stats")

	wakaAPIKey := os.Getenv("INPUT_WAKATIME_API_KEY")
	if wakaAPIKey == "" {
		return fmt.Errorf("Wakatime API Key is required")
	}

	githubToken := os.Getenv("INPUT_GH_TOKEN")
	if githubToken == "" {
		return fmt.Errorf("GitHub Token is required")
	}

	return nil
}
