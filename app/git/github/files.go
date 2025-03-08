package gitGithub

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/pandaci-com/pandaci/types"
	typesDB "github.com/pandaci-com/pandaci/types/database"
	"github.com/rs/zerolog/log"
)

type gitFile struct {
	hash string
	path string
}

func buildFileContentQuery(owner string, repo string, sha string, gitFiles []gitFile) *string {
	query := fmt.Sprintf(`query {
    repository(owner: "%s", name: "%s") {
  `, owner, repo)

	for _, file := range gitFiles {
		query += fmt.Sprintf(`_%s: object(expression: "%s") {
      ... on Blob {
        text
      }
    }`, file.hash, file.hash)
	}

	query += "}}"

	return &query
}

func (c *GithubInstallationClient) GetWorkflowFiles(ctx context.Context, project typesDB.Project, event types.TriggerEvent) ([]types.WorkflowFile, error) {

	repoID, err := strconv.Atoi(project.GitProviderRepoID)
	if err != nil {
		return nil, err
	}

	repo, resp, err := c.Client.Repositories.GetByID(ctx, int64(repoID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return nil, fmt.Errorf("unable to get github repository, failed with StatusCode: %d ", resp.StatusCode)
	}

	tree, resp, err := c.Client.Git.GetTree(ctx, repo.GetOwner().GetLogin(), repo.GetName(), fmt.Sprintf("%s:.pandaci/", event.SHA), true)
	if err != nil && (resp == nil || resp.StatusCode != 404) {
		return nil, err
	} else if resp.StatusCode == 404 {
		return []types.WorkflowFile{}, nil
	}

	re := regexp.MustCompile(`.*\.workflow\.(ts|js|cjs|mjs)$`)

	var gitFiles []gitFile
	for _, entry := range tree.Entries {
		if entry.GetType() == "blob" && re.MatchString(entry.GetPath()) {
			gitFiles = append(gitFiles, gitFile{
				hash: entry.GetSHA(),
				path: entry.GetPath(),
			})
		}
	}

	query := buildFileContentQuery(repo.GetOwner().GetLogin(), repo.GetName(), event.SHA, gitFiles)

	type response struct {
		Data struct {
			Repository map[string]struct {
				Text string `json:"text"`
			} `json:"repository"`
		} `json:"data"`
	}

	data := &response{}
	if err := c.graphqlClient.Query(ctx, data, *query, nil); err != nil {
		return nil, err
	}

	var workflowFiles []types.WorkflowFile

	for _, file := range gitFiles {
		content := data.Data.Repository["_"+file.hash].Text

		if content == "" {

			log.Debug().Msgf("Skipping file: %s, with git hash: %s", file.path, file.hash)
			log.Info().Msgf("Skipping file either due to it being empty or being unable to retreive it's contents")
			continue
		}

		workflowFiles = append(workflowFiles, types.WorkflowFile{
			Hash:    file.hash,
			Content: content,
			Path:    filepath.Join(".pandaci/", file.path),
		})
	}

	return workflowFiles, nil
}
