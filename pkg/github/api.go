package github

import (
	"context"

	"github.com/google/go-github/v65/github"
	"github.com/user/wakatime-profile-stats/internal/errors"
)

const (
	PerPageCount = 10
)

type Client struct {
	client *github.Client
}

func NewGithubClient(authToken string) (*Client, error) {
	if authToken == "" {
		return nil, errors.ErrGithubTokenRequired
	}

	return &Client{client: github.NewClient(nil).WithAuthToken(authToken)}, nil
}

func (gh *Client) GetUser() (*github.User, error) {
	user, _, err := gh.client.Users.Get(context.Background(), "")
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (gh *Client) GetRepos() ([]*github.Repository, error) {
	var allRepos []*github.Repository
	opt := &github.RepositoryListByAuthenticatedUserOptions{
		ListOptions: github.ListOptions{PerPage: PerPageCount},
	}

	for {
		repos, resp, err := gh.client.Repositories.ListByAuthenticatedUser(context.Background(), opt)
		if err != nil {
			return nil, err
		}

		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}

	return allRepos, nil
}
