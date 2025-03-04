package queriesWorkflow

import (
	"context"
	"time"
)

func (q *WorkflowQueries) GetBuildMinutes(ctx context.Context, orgID string, startDate time.Time, endDate time.Time) (int, error) {

	query := `SELECT
			COALESCE(SUM(workflow_run.build_minutes), 0) workflow_mins,
			COALESCE(SUM(job_run.build_minutes), 0) as job_mins
		FROM project
			JOIN workflow_run ON workflow_run.project_id = project.id
			JOIN job_run ON job_run.workflow_run_id = workflow_run.id
		WHERE workflow_run.created_at >= $1 AND workflow_run.created_at < $2 AND project.org_id = $3`

	type Data struct {
		WorkflowMins int `db:"workflow_mins"`
		JobMins      int `db:"job_mins"`
	}

	var data Data
	err := q.GetContext(ctx, &data, query, startDate, endDate, orgID)
	if err != nil {
		return 0, err
	}

	return data.WorkflowMins + data.JobMins, nil
}

type Committer struct {
	UserID         *string `db:"user_id"`
	CommitterEmail *string `db:"committer_email"`
}

func (q *WorkflowQueries) CountCommitters(ctx context.Context, orgID string, startDate time.Time, endDate time.Time, new *[]Committer) (int, error) {
	query := `SELECT DISTINCT
		user_id,
		committer_email
	FROM workflow_run
	JOIN project ON project.id = workflow_run.project_id
	WHERE workflow_run.created_at >= $1 AND workflow_run.created_at < $2 AND project.org_id = $3 AND workflow_run.conclusion != 'failure'`
	// TODO - we need to introduce a new workflow run conclusion called fatal where a system error has occured, not a workflow error.
	// then we should updated the query above to exclude fatal errors from the count, not failures.

	var data []Committer
	err := q.SelectContext(ctx, &data, query, startDate, endDate, orgID)
	if err != nil {
		return 0, err
	}

	if new != nil {
		data = append(data, *new...)
	}

	count := 0

	emailHash := make(map[string]bool)
	userIDHash := make(map[string]bool)

	// Count distinct users
	for _, d := range data {
		if d.UserID != nil {
			if _, ok := userIDHash[*d.UserID]; !ok {
				userIDHash[*d.UserID] = true
				count++
			}

			if d.CommitterEmail != nil {
				emailHash[*d.CommitterEmail] = true
			}
		}
	}

	// Count emails that aren't attached to a user
	for _, d := range data {
		if d.CommitterEmail != nil && d.UserID == nil {
			if _, ok := emailHash[*d.CommitterEmail]; !ok {
				emailHash[*d.CommitterEmail] = true
				count++
			}
		}
	}

	return count, nil
}
