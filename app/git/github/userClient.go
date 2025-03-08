package gitGithub

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/google/go-github/v68/github"
	gitShared "github.com/pandaci-com/pandaci/app/git/shared"
	"github.com/pandaci-com/pandaci/pkg/utils/env"
	"github.com/pandaci-com/pandaci/platform/identity"
	"github.com/pandaci-com/pandaci/types"
	typesDB "github.com/pandaci-com/pandaci/types/database"
	"github.com/rs/zerolog/log"
)

func (c *GithubClient) getUserRedirect(ctx context.Context, user types.User, account *typesDB.UserAccount) (string, error) {
	refreshState, err := c.Queries.CreateAccountRefreshState(ctx, account)
	if err != nil {
		return "", err
	}

	clientID, err := env.GetGithubAppClientID()
	if err != nil {
		return "", err
	}

	backendUrl, err := env.GetBackendURL()
	if err != nil {
		return "", err
	}

	redirectURL := *backendUrl + "/v1/git/github/callback"

	// TODO - we should store the users github username so we can use it to prefill the login field

	return fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&state=%s&allow_signup=false", *clientID, url.QueryEscape(redirectURL), refreshState.ID), nil
}

func (c *GithubClient) RefreshOAuthTokens(ctx context.Context, user types.User, account *typesDB.UserAccount, refreshToken *string, code *string) error {
	if account.UserID == "" {
		account.UserID = user.ID
	} else if account.UserID != user.ID {
		return fmt.Errorf("account userID doesn't match users ID")
	}

	if account.Type == "" {
		account.Type = typesDB.UserAccountTypeGithub
	} else if account.UserID != user.ID {
		return fmt.Errorf("using account type: %s inside github handler", account.Type)
	}

	log.Debug().Msgf("refreshing github token refreshToken")

	clientID, err := env.GetGithubAppClientID()
	if err != nil {
		return err
	}

	clientSecret, err := env.GetGithubAppClientSecret()
	if err != nil {
		return err
	}

	var refreshURL string
	if refreshToken != nil {
		refreshURL = fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&grant_type=refresh_token&refresh_token=%s", *clientID, *clientSecret, *refreshToken)
	} else if code != nil {

		backendUrl, err := env.GetBackendURL()
		if err != nil {
			return err
		}

		redirectURL := *backendUrl + "/v1/git/github/callback"

		refreshURL = fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&redirect_uri=%s&code=%s", *clientID, *clientSecret, redirectURL, *code)
	} else {
		return fmt.Errorf("refresh token or code is required")
	}

	req, err := http.NewRequest("GET", refreshURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 400 {
		log.Error().Msg("Our refresh token is invalid, it's probably been superseded")

		url, err := c.getUserRedirect(ctx, user, account)
		if err != nil {
			return err
		}

		return &gitShared.GitOAuthError{
			GitType:     typesDB.UserAccountTypeGithub,
			Type:        gitShared.GitOAuthErrorTypeExpiredRefreshToken,
			Message:     "refresh token has expired, please reauthenticate",
			RedirectURL: url,
		}
	} else if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return fmt.Errorf("Github oauth failed with error code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type githubRefreshTokenResponse struct {
		AccessToken           string `json:"access_token"`
		ExpiresIn             int    `json:"expires_in"`
		RefreshToken          string `json:"refresh_token"`
		RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
		TokenType             string `json:"token_type"`
		Scope                 string `json:"scope"`
	}

	response := githubRefreshTokenResponse{}
	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	account.AccessToken = response.AccessToken
	account.AccessTokenExpiresAt = time.Now().Add(time.Second * time.Duration(response.ExpiresIn))
	account.RefreshToken = response.RefreshToken
	account.RefreshTokenExpiresAt = time.Now().Add(time.Second * time.Duration(response.RefreshTokenExpiresIn))

	if account.ProviderAccountID == "" {

		provider, err := identity.GetGithubProvider(ctx, user.ID)
		if err != nil {
			return err
		} else if provider == nil {
			return &gitShared.GitOAuthError{
				GitType: typesDB.UserAccountTypeGithub,
				Type:    gitShared.GitOAuthErrorTypeMismatchedAccount,
				Message: "Github account isn't linked, please link your account via the settings to continue",
			}
		}

		client := github.NewClient(nil).WithAuthToken(account.AccessToken)

		githubUser, _, err := client.Users.Get(ctx, "")
		if err != nil {
			return err
		}

		if strconv.Itoa(int(githubUser.GetID())) != provider.ProviderAccountID {
			return &gitShared.GitOAuthError{
				GitType: typesDB.UserAccountTypeGithub,
				Type:    gitShared.GitOAuthErrorTypeMismatchedAccount,
				Message: "Github account is different to the linked account. Please sign in with the correct account or update your linked accounts via the settings to continue",
			}
		}

		account.ProviderAccountID = strconv.Itoa(int(githubUser.GetID()))
	}

	return c.Queries.UpsertAccount(ctx, account)
}

func (c *GithubClient) NewUserClient(ctx context.Context, user types.User) (gitShared.UserClient, error) {
	account, err := c.Queries.GetUserAccountByType(ctx, user, typesDB.UserAccountTypeGithub)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {

		// If a user has linked a github account or signed in with one, then ory has a refresh token we can use
		provider, err := identity.GetGithubProvider(ctx, user.ID)
		if err != nil {
			return nil, err
		} else if provider == nil {
			return nil, fmt.Errorf("No Github providers available")
		}

		account.UserID = user.ID
		account.Type = typesDB.UserAccountTypeGithub
		account.ProviderAccountID = provider.ProviderAccountID

		if err := c.RefreshOAuthTokens(ctx, user, account, &provider.InitialRefreshToken, nil); err != nil {
			return nil, err
		}

		// We subtract a minute from the expiry time to make sure we don't run into any issues with the token expiring before we use it
	} else if account.AccessTokenExpiresAt.Before(time.Now().Add(-time.Minute)) {
		log.Info().Msgf("Github access token has expired for user %s", user.ID)

		if account.RefreshTokenExpiresAt.Before(time.Now().Add(-time.Minute)) {
			log.Info().Msgf("Github refresh token has expired for user %s", user.ID)

			url, err := c.getUserRedirect(ctx, user, account)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get user redirect")
				return nil, err
			}

			return nil, &gitShared.GitOAuthError{
				GitType:     typesDB.UserAccountTypeGithub,
				Type:        gitShared.GitOAuthErrorTypeExpiredRefreshToken,
				Message:     "refresh token has expired, please reauthenticate",
				RedirectURL: url,
			}
		}

		if err := c.RefreshOAuthTokens(ctx, user, account, &account.RefreshToken, nil); err != nil {
			return nil, err
		}
	}

	badTokenTripper := &BadTokenTripper{
		Transport: http.DefaultTransport,
		GetRedirectURL: func() (string, error) {
			return c.getUserRedirect(ctx, user, account)
		},
	}

	httpClient := &http.Client{
		Transport: badTokenTripper,
	}

	client := github.NewClient(httpClient).WithAuthToken(account.AccessToken)

	return &GithubUserClient{
		queries:      c.Queries,
		githubClient: client,
		gitClient:    c,
	}, nil
}

type BadTokenTripper struct {
	Transport      http.RoundTripper
	GetRedirectURL func() (string, error)
}

// Catches any 401 requets back from github so we can attempt to get the user to reauthenticate
func (crt *BadTokenTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := crt.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		log.Error().Msg("A request to github has been made with a bad token")

		url, err := crt.GetRedirectURL()
		if err != nil {
			return nil, err
		}

		return nil, &gitShared.GitOAuthError{
			GitType:     typesDB.UserAccountTypeGithub,
			Type:        gitShared.GitOAuthErrorTypeExpiredRefreshToken,
			Message:     "refresh token has expired, please reauthenticate",
			RedirectURL: url,
		}
	}
	return resp, nil
}
