package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v65/github"
)

type GithubClient struct {
	client *github.Client
}

func NewGithubClient(authToken string) (*GithubClient, error) {
	if authToken == "" {
		return nil, fmt.Errorf("GitHub Token is required")
	}

	return &GithubClient{client: github.NewClient(nil).WithAuthToken(authToken)}, nil
}

func (gh *GithubClient) GetUser() (*github.User, error) {
	user, _, err := gh.client.Users.Get(context.Background(), "")
	if err != nil {
		return nil, fmt.Errorf("getting user: %w", err)
	}

	return user, nil
}

func (gh *GithubClient) GetRepos() ([]*github.Repository, error) {
	var allRepos []*github.Repository
	opt := &github.RepositoryListByAuthenticatedUserOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	for {
		repos, resp, err := gh.client.Repositories.ListByAuthenticatedUser(context.Background(), opt)
		if err != nil {
			return nil, fmt.Errorf("getting repos: %w", err)
		}

		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}

	return allRepos, nil
}
