package handlersProject

import (
	"fmt"
	"net/http"
	"os"
	"slices"

	"connectrpc.com/connect"
	"github.com/labstack/echo/v4"
	middlewareLoaders "github.com/pandaci-com/pandaci/app/api/middleware/loaders"
	grpcMiddleware "github.com/pandaci-com/pandaci/app/grpc/middleware"
	"github.com/pandaci-com/pandaci/pkg/jwt"
	"github.com/pandaci-com/pandaci/pkg/retryClient"
	"github.com/pandaci-com/pandaci/pkg/utils"
	"github.com/pandaci-com/pandaci/pkg/utils/env"
	"github.com/pandaci-com/pandaci/platform/identity"
	pb "github.com/pandaci-com/pandaci/proto/go/v1"
	pbConnect "github.com/pandaci-com/pandaci/proto/go/v1/v1connect"
	"github.com/pandaci-com/pandaci/types"
	typesHTTP "github.com/pandaci-com/pandaci/types/http"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetLogStream(c echo.Context) error {
	project, err := middlewareLoaders.GetProject(c)
	if err != nil {
		return err
	}

	workflowRun, err := middlewareLoaders.GetWorkflowRun(c)
	if err != nil {
		return err
	}

	stepID := c.QueryParam("step_id")

	// TODO - have a way to get this from the runner name in the workflow
	runnerAddress, err := env.GetRunnerAddress()
	if err != nil {
		return err
	}

	runnerClient := pbConnect.NewRunnerServiceClient(&http.Client{
		Transport: &retryClient.RetryRoundTripper{
			Base: http.DefaultTransport,
		},
	}, *runnerAddress, connect.WithInterceptors(grpcMiddleware.NewWorkflowJWTInerceptor(*h.jwtHandler, nil)))

	jwt, err := h.jwtHandler.CreateWorkflowToken(jwt.WorkflowClaims{
		WorkflowID: workflowRun.ID,
		OrgID:      project.OrgID,
		ProjectID:  project.ID,
	})
	if err != nil {
		log.Error().Err(err).Msg("creating workflow token")
		return err
	}

	streamURL, err := runnerClient.GetLogStream(c.Request().Context(), connect.NewRequest(&pb.RunnerServiceGetLogStreamRequest{
		WorkflowJwt: jwt,
		StepId:      &stepID,
	}))
	if err != nil {
		log.Error().Err(err).Msg("getting log stream")
		return err
	}

	res := typesHTTP.LogStream{
		URL:           streamURL.Msg.Url,
		Authorization: fmt.Sprintf("Bearer %s", jwt),
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) GetWorkflowRunLogs(c echo.Context) error {
	project, err := middlewareLoaders.GetProject(c)
	if err != nil {
		return err
	}

	workflowRun, err := middlewareLoaders.GetWorkflowRun(c)
	if err != nil {
		return err
	}

	logID := c.Param("log_id")

	url, err := h.bucketClient.GetObject(c.Request().Context(), os.Getenv("WORKFLOW_LOGS_BUCKET"), fmt.Sprintf("%s/%s/%s/%s.csv", project.OrgID, project.ID, workflowRun.ID, logID), 60*60*24*7)
	if err != nil {
		log.Error().Err(err).Msg("getting step presigned url")
		return err
	}

	return c.JSON(http.StatusOK, typesHTTP.Log{URL: url.URL})
}

func (h *Handler) GetWorkflowRunWithItems(c echo.Context) error {
	workflowRun, err := middlewareLoaders.GetWorkflowRun(c)
	if err != nil {
		return err
	}

	project, err := middlewareLoaders.GetProject(c)
	if err != nil {
		return err
	}

	commitURL := make(chan string)
	go func() {
		provider, err := h.queries.GetGitIntegration(c.Request().Context(), project.GitIntegrationID)
		if err != nil {
			log.Error().Err(err).Msg("getting provider")
			return
		}

		gitClient, err := h.gitHandler.GetClient(provider.Type)
		if err != nil {
			log.Error().Err(err).Msg("getting git client")
			return
		}

		installationClient, err := gitClient.NewInstallationClient(c.Request().Context(), provider.ProviderID)
		if err != nil {
			log.Error().Err(err).Msg("getting installation client")
			return
		}

		repo, err := installationClient.GetProjectGitRepoData(c.Request().Context(), *project)
		if err != nil {
			log.Error().Err(err).Msg("getting repo")
			return
		}

		commitURL <- gitClient.NewInstallationClient(c.Request().Context(), project.Git).GetCommitURL(c.Request().Context(), project.RepoID, workflowRun.GitSha)

		// h.gitHandler.GetClient(project.)
	}()

	jobs, err := h.queries.GetWorkflowRunJobsHTTP(c.Request().Context(), workflowRun)
	if err != nil {
		return err
	}

	if jobs == nil {
		jobs = []typesHTTP.JobRun{}
	}

	alerts := workflowRun.GetAlerts()
	if err != nil {
		log.Error().Err(err).Msg("getting alerts")
		return err
	}

	committer := typesHTTP.Committer{
		Email: *workflowRun.CommitterEmail,
	}

	if workflowRun.UserID != nil {
		user, err := identity.GetUserByID(c.Request().Context(), *workflowRun.UserID)
		if err != nil {
			return err
		}
		if user.Name != nil {
			committer.Name = *user.Name
		}
		if user.Avatar != nil {
			committer.Avatar = *user.Avatar
		}
	}

	return c.JSON(http.StatusOK, typesHTTP.WorkflowRun{
		ID:         workflowRun.ID,
		Number:     workflowRun.Number,
		CreatedAt:  workflowRun.CreatedAt,
		FinishedAt: workflowRun.FinishedAt,
		ProjectID:  workflowRun.ProjectID,
		Status:     workflowRun.Status,
		Conclusion: workflowRun.Conclusion,
		Name:       workflowRun.Name,
		GitSha:     workflowRun.GitSha,
		GitBranch:  workflowRun.GitBranch,
		Jobs:       jobs,
		Alerts:     alerts,
		PrNumber:   workflowRun.PrNumber,
		Trigger:    workflowRun.Trigger,
		GitTitle:   workflowRun.GitTitle,
		Committer:  committer,
	})
}

func (h *Handler) GetWorkflowRuns(c echo.Context) error {
	project, err := middlewareLoaders.GetProject(c)
	if err != nil {
		return err
	}

	page, err := utils.GetQueryParamInt(c, "page")
	if err != nil {
		return err
	}
	if page == nil {
		page = types.Pointer(0)
	}

	perPage, err := utils.GetQueryParamInt(c, "per_page")
	if err != nil {
		return err
	}
	if perPage == nil {
		perPage = types.Pointer(25)
	}

	workflowsDB, hasMore, err := h.queries.GetProjectWorkflows(c.Request().Context(), project, *page, *perPage)
	if err != nil {
		return err
	}

	userIDs := []string{}
	for _, workflow := range *workflowsDB {
		if workflow.UserID != nil && !slices.Contains(userIDs, *workflow.UserID) {
			userIDs = append(userIDs, *workflow.UserID)
		}
	}

	fmt.Println("userIDs", userIDs)

	oryUsers, err := identity.ListUsersByIDs(c.Request().Context(), userIDs)
	if err != nil {
		return err
	}

	workflowsHTTP := make([]typesHTTP.WorkflowRun, len(*workflowsDB))
	for i, workflow := range *workflowsDB {

		committer := typesHTTP.Committer{}

		if workflow.CommitterEmail != nil {
			committer.Email = *workflow.CommitterEmail
		}

		if workflow.UserID != nil {
			userIndex := slices.IndexFunc(oryUsers, func(user *types.User) bool {
				return user.ID == *workflow.UserID
			})

			if userIndex != -1 {
				committer.Name = *oryUsers[userIndex].Name
				committer.Avatar = *oryUsers[userIndex].Avatar
				if committer.Email == "" {
					committer.Email = oryUsers[userIndex].Email
				}
			}
		}

		workflowsHTTP[i] = typesHTTP.WorkflowRun{
			ID:         workflow.ID,
			ProjectID:  workflow.ProjectID,
			Name:       workflow.Name,
			CreatedAt:  workflow.CreatedAt,
			Status:     workflow.Status,
			Number:     workflow.Number,
			GitSha:     workflow.GitSha,
			GitBranch:  workflow.GitBranch,
			FinishedAt: workflow.FinishedAt,
			Conclusion: workflow.Conclusion,
			PrNumber:   workflow.PrNumber,
			Trigger:    workflow.Trigger,
			Committer:  committer,
			GitTitle:   workflow.GitTitle,
		}
	}

	return c.JSON(http.StatusOK, typesHTTP.Paginated{
		Data: workflowsHTTP,
		Next: hasMore,
	})
}
