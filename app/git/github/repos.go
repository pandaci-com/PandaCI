package gitGithub

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v68/github"
	gitShared "github.com/pandaci-com/pandaci/app/git/shared"
	"github.com/pandaci-com/pandaci/types"
	typesDB "github.com/pandaci-com/pandaci/types/database"
	typesHTTP "github.com/pandaci-com/pandaci/types/http"
)

func formatRepos(repos []*github.Repository) []typesHTTP.GitRepository {
	formattedRepos := []typesHTTP.GitRepository{}
	for _, repo := range repos {
		formattedRepos = append(formattedRepos, typesHTTP.GitRepository{
			ID:          strconv.Itoa(int(repo.GetID())),
			Name:        repo.GetName(),
			UpdatedAt:   repo.GetUpdatedAt().Time,
			Description: repo.GetDescription(),
			Public:      repo.GetVisibility() == "public",
			URL:         repo.GetHTMLURL(),
		})
	}

	return formattedRepos
}

func (c *GithubUserClient) getRepoByName(ctx context.Context, installationID int, owner string, name string) (*github.Repository, error) {
	installationClient, err := c.gitClient.NewGithubInstallationClient(ctx, strconv.Itoa(installationID))
	if err != nil {
		return nil, err
	}

	repo, resp, err := installationClient.Client.Repositories.Get(ctx, owner, name)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, nil
	}

	return repo, nil
}

func (c *GithubUserClient) GetRepositories(ctx context.Context, installationID string, options gitShared.GetRepositoriesOptions) (*typesHTTP.GitRepositories, error) {
	installID, err := strconv.Atoi(installationID)
	if err != nil {
		return nil, err
	}

	if options.Name != "" {
		repo, err := c.getRepoByName(ctx, installID, options.Owner, options.Name)
		if err != nil {
			return nil, err
		}

		var repos []typesHTTP.GitRepository
		if repo != nil {
			repos = append(repos, typesHTTP.GitRepository{
				ID:          strconv.Itoa(int(repo.GetID())),
				Name:        repo.GetName(),
				UpdatedAt:   repo.GetUpdatedAt().Time,
				Description: repo.GetDescription(),
				URL:         repo.GetHTMLURL(),
				Public:      repo.GetVisibility() == "public",
			})
		}

		return &typesHTTP.GitRepositories{
			Repos:         repos,
			LimitExceeded: false,
			Limit:         100,
		}, nil
	}

	if options.Query == "" {
		// No query so we can just use the dedicated github api
		listRepos, resp, err := c.githubClient.Apps.ListUserRepos(ctx, int64(installID), &github.ListOptions{
			PerPage: 100,
		})
		if err != nil {
			return nil, err
		}

		if resp.NextPage != 0 {
			// Too many scoped repos. Users must search for exact matches or ensure the installation can see all repos
			return &typesHTTP.GitRepositories{
				Limit:         100,
				LimitExceeded: true,
			}, nil
		} else {
			return &typesHTTP.GitRepositories{
				Limit:         100,
				LimitExceeded: false,
				Repos:         formatRepos(listRepos.Repositories),
			}, nil
		}
	}

	// If just passed a query, we run it with the users credentials
	// We need to make sure we are only searching for repos which our app can see (this responsiblity is passed onto the client)
	// this avoids an extra api call
	// We should only be using the query if we know for sure that the installation isn't scoped to any specific repos

	searchUserRepos, _, err := c.githubClient.Search.Repositories(ctx, options.Query, &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 5,
		},
	})
	if err != nil {
		return nil, err
	}

	return &typesHTTP.GitRepositories{
		Repos:         formatRepos(searchUserRepos.Repositories),
		LimitExceeded: false,
		Limit:         100,
	}, nil
}

func (c *GithubInstallationClient) RepoExists(ctx context.Context, repoID string) (bool, error) {
	intID, err := strconv.Atoi(repoID)
	if err != nil {
		return false, err
	}

	_, resp, err := c.Client.Repositories.GetByID(ctx, int64(intID))
	if err != nil {
		return false, err
	}

	if resp.StatusCode == 404 {
		return false, nil
	}

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return false, fmt.Errorf("error whilst fetching github repo, got status code: %d", resp.StatusCode)
	}

	return true, nil
}

func (c *GithubInstallationClient) GetProjectGitRepoData(ctx context.Context, project typesDB.Project) (*types.GitRepoData, error) {
	repoID, err := strconv.Atoi(project.GitProviderRepoID)
	if err != nil {
		return nil, err
	}

	repo, resp, err := c.Client.Repositories.GetByID(ctx, int64(repoID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return nil, fmt.Errorf("error whilst fetching github repo, got status code: %d", resp.StatusCode)
	}

	token, err := c.itr.Token(ctx)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://x-access-token:%s@%s/%s/%s.git", token, "github.com", repo.GetOwner().GetLogin(), repo.GetName())

	return &types.GitRepoData{
		URL:           url,
		DefaultBranch: repo.GetDefaultBranch(),
	}, nil
}
