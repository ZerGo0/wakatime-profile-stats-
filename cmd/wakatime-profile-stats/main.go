package main

import (
	"fmt"
	"log"
	"os"

	"github.com/user/wakatime-profile-stats/pkg/git"
	"github.com/user/wakatime-profile-stats/pkg/github"
	"github.com/user/wakatime-profile-stats/pkg/stats"
	"github.com/user/wakatime-profile-stats/pkg/wakatime"

	go_github "github.com/google/go-github/v65/github"
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

	wakaAPIKey, githubToken, err := retrieveEnvVars()
	if err != nil {
		return fmt.Errorf("new function: %w", err)
	}

	repos, repoPath, err := getGithubRepos(githubToken)
	if err != nil {
		return fmt.Errorf("getting github user and repos: %w", err)
	}

	sevenDaysStats, monthlyStats, yearlyStats, allTimeStats, err := getWakaStats(wakaAPIKey)
	if err != nil {
		return fmt.Errorf("getting wakatime stats: %w", err)
	}

	err = updateReadmeStats(repoPath, sevenDaysStats, monthlyStats, yearlyStats, allTimeStats, repos)
	if err != nil {
		return fmt.Errorf("updating readme stats: %w", err)
	}

	zap.L().Info("Wakatime Profile Stats completed successfully")

	return nil
}

func retrieveEnvVars() (string, string, error) {
	wakaAPIKey := os.Getenv("INPUT_WAKATIME_API_KEY")
	if wakaAPIKey == "" {
		return "", "", fmt.Errorf("GitHub Token is required")
	}

	githubToken := os.Getenv("INPUT_GH_TOKEN")
	if githubToken == "" {
		return "", "", fmt.Errorf("GitHub Token is required")
	}

	zap.L().Info("Environment variables are present and valid")
	return wakaAPIKey, githubToken, nil
}

func getGithubRepos(githubToken string) ([]*go_github.Repository, string, error) {
	gClient, err := github.NewGithubClient(githubToken)
	if err != nil {
		return nil, "", fmt.Errorf("new github client: %w", err)
	}

	user, err := gClient.GetUser()
	if err != nil {
		return nil, "", fmt.Errorf("getting user: %w", err)
	}

	repos, err := gClient.GetRepos()
	if err != nil {
		return nil, "", fmt.Errorf("getting repos: %w", err)
	}

	remoteName := (*user.Login) + "/" + (*user.Login)
	repoPath := "https://" + githubToken + "@github.com/" + remoteName + ".git"

	zap.L().Info("Github login was successful", zap.Int("repos", len(repos)))

	return repos, repoPath, nil
}

func getWakaStats(wakaAPIKey string) (*wakatime.WakaStats, *wakatime.WakaStats, *wakatime.WakaStats, *wakatime.WakaStats, error) {
	wClient := wakatime.NewClient("https://wakatime.com/api/v1", wakaAPIKey)
	sevenDaysStats, err := wClient.GetStats("last_7_days")
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("getting stats: %w", err)
	}

	monthlyStats, err := wClient.GetStats("last_30_days")
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("getting stats: %w", err)
	}

	yearlyStats, err := wClient.GetStats("last_year")
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("getting stats: %w", err)
	}

	allTimeStats, err := wClient.GetStats("all_time")
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("getting stats: %w", err)
	}

	zap.L().Info("Got Wakatime stats")
	return sevenDaysStats, monthlyStats, yearlyStats, allTimeStats, nil
}

func updateReadmeStats(repoPath string, sevenDaysStats, monthlyStats, yearlyStats, allTimeStats *wakatime.WakaStats, githubRepos []*go_github.Repository) error {
	g, err := git.SetupRepo(repoPath)
	if err != nil {
		return fmt.Errorf("setting up repo: %w", err)
	}

	textStats, err := stats.ProcessStats(sevenDaysStats, monthlyStats, yearlyStats, allTimeStats, githubRepos)
	if err != nil {
		return fmt.Errorf("processing stats: %w", err)
	}

	err = g.UpdateStats(*textStats)
	if err != nil {
		return fmt.Errorf("updating stats: %w", err)
	}

	err = g.CommitAndPush()
	if err != nil {
		return fmt.Errorf("committing and pushing: %w", err)
	}

	return nil
}
