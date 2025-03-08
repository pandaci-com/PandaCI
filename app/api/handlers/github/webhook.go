package handlersGithub

import (
	"context"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/go-playground/webhooks/v6/github"
	ghAPI "github.com/google/go-github/v68/github"
	"github.com/labstack/echo/v4"
	gitGithub "github.com/pandaci-com/pandaci/app/git/github"
	"github.com/pandaci-com/pandaci/pkg/utils/env"
	"github.com/pandaci-com/pandaci/types"
	typesDB "github.com/pandaci-com/pandaci/types/database"
	"github.com/rs/zerolog/log"
)

func (h *Handler) getCommitter(ctx context.Context, installID int64, email string, userName string) types.Committer {

	if userName == "" {
		return types.Committer{
			Email: &email,
		}
	}

	gh, err := h.githubClient.NewInstallationClient(ctx, strconv.FormatInt(installID, 10))
	if err != nil {
		log.Error().Err(err).Msg("Failed to get github client")
		return types.Committer{
			Email: &email,
		}
	}

	githubID, err := gh.GetUserIDFromUsername(ctx, userName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get github user id")
		return types.Committer{
			Email: &email,
		}
	}

	account, err := h.queries.GetUserAccountByProviderAccountID(ctx, strconv.FormatInt(githubID, 10), typesDB.UserAccountTypeGithub)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user account")
		return types.Committer{
			Email: &email,
		}
	}

	return types.Committer{
		Email:  &email,
		UserID: &account.UserID,
	}
}

func (h *Handler) startProjects(triggerEvent types.TriggerEvent, projects []typesDB.Project) {

	wg := sync.WaitGroup{}
	wg.Add(len(projects))

	for _, project := range projects {

		go func(project typesDB.Project, ctx context.Context) {

			defer wg.Done()

			workflows, err := h.scanner.GetWorkflowDefinitions(ctx, project, triggerEvent)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get workflow definitions")
				return
			}

			log.Debug().Interface("workflows", workflows).Msg("Got workflows")

			runsDB, err := h.orchestrator.StartWorkflows(ctx, &project, workflows)
			if err != nil {
				log.Error().Err(err).Msg("Failed to start workflows")
				return
			}

			log.Debug().Interface("runs", runsDB).Msg("Started workflows")
		}(project, context.Background())

	}

	wg.Wait()
}

func (h *Handler) handlePushEvent(ctx context.Context, payload github.PushPayload) {
	log.Debug().Interface("payload", payload).Msg("Received github push event")

	gitProvider, err := h.queries.GetGitIntegrationByProviderID(ctx, strconv.FormatInt(int64(payload.Installation.ID), 10), types.GitProviderTypeGithub)
	if err != nil {
		// TODO - we should trigger a sync of the git provider
		log.Error().Err(err).Msg("Git provider not found")
		return
	}

	projects, err := h.queries.GetProjectsByGitIntegrationID(ctx, gitProvider.ID, strconv.FormatInt(payload.Repository.ID, 10))
	if err != nil {
		log.Error().Err(err).Msg("Project not found")
		return
	}

	if !strings.HasPrefix(payload.Ref, "refs/heads/") {
		log.Error().Msg("Invalid ref")
		return
	}

	branch := strings.TrimPrefix(payload.Ref, "refs/heads/")

	email := payload.HeadCommit.Author.Email
	username := payload.HeadCommit.Author.Username

	triggerEvent := types.TriggerEvent{
		Source:    types.TriggerEventSourceGithub,
		SHA:       payload.After,
		GitTitle:  payload.HeadCommit.Message,
		Branch:    branch,
		Trigger:   types.RunTriggerPush,
		Committer: h.getCommitter(ctx, int64(payload.Installation.ID), email, username),
	}

	h.startProjects(triggerEvent, *projects)
}

func (h *Handler) handlePullRequestEvent(ctx context.Context, payload github.PullRequestPayload) {
	log.Debug().Interface("payload", payload).Msg("Received github push event")

	if !slices.Contains([]string{"opened", "closed", "synchronize", "reopened"}, payload.Action) {
		log.Debug().Msg("Ignoring pull request action")
		return
	}

	gitProvider, err := h.queries.GetGitIntegrationByProviderID(ctx, strconv.FormatInt(int64(payload.Installation.ID), 10), types.GitProviderTypeGithub)
	if err != nil {
		// TODO - we should trigger a sync of the git provider
		log.Error().Err(err).Msg("Git provider not found")
		return
	}

	if payload.PullRequest.Head.Repo.ID != payload.Repository.ID {
		log.Error().Msg("Forks are not currently supported")

		ghClient := gitGithub.GithubClient{
			Queries: h.queries,
		}

		installClient, err := ghClient.NewGithubInstallationClient(ctx, strconv.FormatInt(int64(payload.Installation.ID), 10))
		if err != nil {
			log.Error().Err(err).Msg("Failed to get github installation client")
			return
		}

		if _, _, err := installClient.Client.Repositories.CreateStatus(ctx, payload.Repository.Owner.Login, payload.Repository.Name, payload.PullRequest.Head.Sha, &ghAPI.RepoStatus{
			State:       types.Pointer("failure"),
			Description: types.Pointer("Pull requests from forks are not currently supported"),
			Context:     types.Pointer("PandaCI"),
			TargetURL:   types.Pointer("https://pandaci.com/docs/platform/workflows/env#git-fork-protection"),
		}); err != nil {
			log.Error().Err(err).Msg("Failed to create status")
		}

		return
	}

	projects, err := h.queries.GetProjectsByGitIntegrationID(ctx, gitProvider.ID, strconv.FormatInt(payload.Repository.ID, 10))
	if err != nil {
		log.Error().Err(err).Msg("Project not found")
		return
	}

	trigger := types.RunTriggerPullRequestOpened
	if payload.Action == "closed" {
		trigger = types.RunTriggerPullRequestClosed
	} else if payload.Action == "synchronize" {
		trigger = types.RunTriggerPullRequestSynchronize
	}

	triggerEvent := types.TriggerEvent{
		Source:       types.TriggerEventSourceGithub,
		SHA:          payload.PullRequest.Head.Sha,
		Branch:       payload.PullRequest.Head.Ref,
		Trigger:      trigger,
		GitTitle:     payload.PullRequest.Title,
		TargetBranch: &payload.PullRequest.Base.Ref,
		PrNumber:     types.Pointer(int32(payload.PullRequest.Number)),
		Committer:    h.getCommitter(ctx, payload.Installation.ID, payload.PullRequest.User.Login, payload.PullRequest.User.Login),
	}

	h.startProjects(triggerEvent, *projects)
}

func (h *Handler) GetGithubWebhook() (echo.HandlerFunc, error) {
	secret, err := env.GetGithubWebhookSecret()
	if err != nil {
		return nil, err
	}

	hook, err := github.New(github.Options.Secret(*secret))
	if err != nil {
		return nil, err
	}
	return func(c echo.Context) error {
		payload, err := hook.Parse(c.Request(), github.PushEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				return c.String(http.StatusNotFound, "event not found")
			}
			log.Error().Err(err).Msg("Failed to parse github webhook payload")
			return err
		}

		switch payload.(type) {
		case github.PushPayload:
			push := payload.(github.PushPayload)
			go h.handlePushEvent(context.Background(), push)

			return c.String(http.StatusOK, "ok")
		case github.PullRequestPayload:
			pr := payload.(github.PullRequestPayload)
			log.Info().Interface("pr", pr).Msg("Received pull request event")
			go h.handlePullRequestEvent(context.Background(), pr)

			return c.String(http.StatusOK, "ok")
		}

		return c.String(http.StatusOK, "ok")
	}, nil
}
