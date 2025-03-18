package gitShared

import (
	"context"
	"fmt"

	"github.com/pandaci-com/pandaci/types"
	typesDB "github.com/pandaci-com/pandaci/types/database"
	typesHTTP "github.com/pandaci-com/pandaci/types/http"
)

type Client interface {
	NewUserClient(ctx context.Context, user types.User) (UserClient, error)
	NewAppClient(ctx context.Context) (AppClient, error)
	NewInstallationClient(ctx context.Context, installationID string) (InstallationClient, error)
	RefreshOAuthTokens(ctx context.Context, user types.User, account *typesDB.UserAccount, refreshToken *string, code *string) error
}

type GetRepositoriesOptions struct {
	Query string
	Name  string
	Owner string
}

type GetInstallationsOptions struct {
	PerPage *int
	Page    *int
}

type UserClient interface {
	GetInstallations(ctx context.Context, options GetInstallationsOptions) (*typesHTTP.GitInstallations, error)
	GetRepositories(ctx context.Context, installationID string, options GetRepositoriesOptions) (*typesHTTP.GitRepositories, error)
}

type AppClient interface {
	GetInstallation(ctx context.Context, installationID string) (*typesHTTP.GitInstallation, error)
}

type InstallationClient interface {
	GetWorkflowFiles(ctx context.Context, project typesDB.Project, event types.TriggerEvent) ([]types.WorkflowFile, error)
	GetProjectGitRepoData(ctx context.Context, projecct typesDB.Project) (*types.GitRepoData, error)
	UpdateRunInRepo(ctx context.Context, org typesDB.OrgDB, project typesDB.Project, run typesDB.WorkflowRun) error
	RepoExists(ctx context.Context, repoID string) (bool, error)
	GetUserIDFromUsername(ctx context.Context, username string) (int64, error)
	PostAwaitingApprovalInRepo(ctx context.Context, org typesDB.OrgDB, project typesDB.Project, run typesDB.WorkflowRun) error
}

type GitOAuthErrorType string

type GitOAuthError struct {
	GitType     typesDB.UserAccountType `json:"gitType"`
	Type        GitOAuthErrorType       `json:"type"`
	RedirectURL string                  `json:"redirectURL"`
	Message     string                  `json:"message"`
}

const (
	GitOAuthErrorTypeExpiredRefreshToken = "ExpiredRefreshToken"
	GitOAuthErrorTypeMismatchedAccount   = "MismatchedAccount"
)

func (e *GitOAuthError) Error() string {
	return fmt.Sprintf("GitType: %s, Type: %s, Msg: %s", e.GitType, e.Type, e.Message)
}
