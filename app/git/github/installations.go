package gitGithub

import (
	"context"
	"net/http"
	"strconv"

	"github.com/google/go-github/v68/github"
	gitShared "github.com/pandaci-com/pandaci/app/git/shared"
	"github.com/pandaci-com/pandaci/types"
	typesHTTP "github.com/pandaci-com/pandaci/types/http"
)

func (c *GithubAppClient) GetInstallation(ctx context.Context, installationID string) (*typesHTTP.GitInstallation, error) {
	installID, err := strconv.Atoi(installationID)
	if err != nil {
		return nil, err
	}

	installation, resp, err := c.githubClient.Apps.GetInstallation(ctx, int64(installID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	formattedInstall := typesHTTP.GitInstallation{
		ID:               installationID,
		Name:             installation.GetAccount().GetLogin(),
		IsUser:           installation.GetAccount().GetType() == "User",
		RepositoryScopes: installation.GetRepositorySelection(),
		AvatarURL:        installation.GetAccount().GetAvatarURL(),
		Type:             types.GitProviderTypeGithub,
		AccountID:        strconv.Itoa(int(installation.GetAccount().GetID())),
	}

	return &formattedInstall, nil
}

func (c *GithubInstallationClient) GetUserIDFromUsername(ctx context.Context, username string) (int64, error) {
	user, _, err := c.Client.Users.Get(ctx, username)
	if err != nil {
		return 0, err
	}

	return user.GetID(), nil
}

func (c *GithubUserClient) GetInstallations(ctx context.Context, options gitShared.GetInstallationsOptions) (*typesHTTP.GitInstallations, error) {
	page := 1
	if options.Page != nil {
		page = *options.Page
	}

	githubOpts := &github.ListOptions{
		Page: page,
	}

	if options.PerPage != nil {
		githubOpts.PerPage = *options.PerPage
	}

	installations, res, err := c.githubClient.Apps.ListUserInstallations(ctx, githubOpts)
	if err != nil {
		return nil, err
	}

	formattedInstalls := make([]typesHTTP.GitInstallation, len(installations))

	for i, install := range installations {
		formattedInstalls[i] = typesHTTP.GitInstallation{
			ID:               strconv.Itoa(int(install.GetID())),
			Name:             install.GetAccount().GetLogin(),
			IsUser:           install.GetAccount().GetType() == "User",
			RepositoryScopes: install.GetRepositorySelection(),
			AvatarURL:        install.GetAccount().GetAvatarURL(),
			Type:             types.GitProviderTypeGithub,
			AccountID:        strconv.Itoa(int(install.GetAccount().GetID())),
		}
	}

	return &typesHTTP.GitInstallations{
		IsLastPage:    res.NextPage == 0,
		Installations: formattedInstalls,
	}, nil
}
