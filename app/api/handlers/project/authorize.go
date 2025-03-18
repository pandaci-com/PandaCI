package handlersProject

import "github.com/labstack/echo/v4"

func (h *Handler) GetRunAuthorization(c echo.Context) error {

	type RunAuthorization struct {
		PullRequestURL string `json:"pullRequestURL"`
		CommitURL      string `json:"commitURL"`
	}

	return nil
}

func (h *Handler) AuthorizeRun(c echo.Context) error {
	return nil
}
