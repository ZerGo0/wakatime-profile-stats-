package errors

import "errors"

var (
	ErrGithubTokenRequired    = errors.New("github token is required")
	ErrWakatimeAPIKeyRequired = errors.New("wakatime api key is required")
	ErrRepoNotSetup           = errors.New("repo is not setup")
	ErrTagNotFound            = errors.New("tag not found")
)
