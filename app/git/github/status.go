package gitGithub

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v68/github"
	"github.com/pandaci-com/pandaci/types"
	typesDB "github.com/pandaci-com/pandaci/types/database"
	"github.com/rs/zerolog/log"
)

func getStatus(status types.RunStatus, conclusion *types.RunConclusion) string {
	// error, success, failure, pending

	if conclusion == nil {
		return "pending"
	}

	if *conclusion == types.RunConclusionSuccess {
		return "success"
	}

	return "failure"
}

func (c *GithubInstallationClient) UpdateRunInRepo(ctx context.Context, org typesDB.OrgDB, project typesDB.Project, run typesDB.WorkflowRun) error {

	if project.GitIntegrationID == "" {
		log.Info().Msg("No GitIntegrationID found, skipping update")
		return nil
	} else if project.GitProviderRepoID == "" {
		log.Info().Msg("No GitProviderRepoID found, skipping update")
		return nil
	}

	status := getStatus(run.Status, run.Conclusion)

	repoID, err := strconv.ParseInt(project.GitProviderRepoID, 10, 64)
	if err != nil {
		return err
	}

	repo, _, err := c.Client.Repositories.GetByID(ctx, repoID)
	if err != nil {
		return err
	}

	var description string
	if run.Conclusion != nil {
		description = fmt.Sprintf("Workflow finished with status: %s", *run.Conclusion)
	} else if run.Status == types.RunStatusQueued {
		description = "PandaCI has queued your workflow"
	} else {
		description = "PandaCI is running your workflow"
	}

	_, _, err = c.Client.Repositories.CreateStatus(ctx, repo.Owner.GetLogin(), repo.GetName(), run.GitSha, &github.RepoStatus{
		TargetURL:   types.Pointer(fmt.Sprintf("https://app.pandaci.com/%s/%s/runs/%d", org.Slug, project.Slug, run.Number)),
		State:       &status,
		Description: types.Pointer(description),
		Context:     types.Pointer(fmt.Sprintf("PandaCI — %s", run.Name)),
	})

	return err
}
