package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/mattermost/mattermost-server/v5/model"
)

// Perform an HTTP POST request to an integration's action endpoint.
// Caller must consume and close returned http.Response as necessary.
// For internal requests, requests are routed directly to a plugin ServerHTTP hook
func (p *Plugin) doActionRequest(rawURL string) (*http.Response, *model.AppError) {
	rawURLPath := path.Clean(rawURL)
	p.API.LogInfo(
		"##################################",
	)
	p.API.LogInfo(
		"rawURLPath",
		"rawURLPath", rawURLPath,
	)
	p.API.LogInfo(
		"##################################",
	)

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return nil, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, err.Error(), http.StatusBadRequest)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Allow access to plugin routes for action buttons
	httpClient := &http.Client{}
	resp, httpErr := httpClient.Do(req)
	if httpErr != nil {
		return nil, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, "err="+httpErr.Error(), http.StatusBadRequest)
	}

	if resp.StatusCode != http.StatusOK {
		return resp, model.NewAppError("DoActionRequest", "api.post.do_action.action_integration.app_error", nil, fmt.Sprintf("status=%v", resp.StatusCode), http.StatusBadRequest)
	}
	p.API.LogInfo(
		"response",
		"resp.StatusCode", resp.StatusCode,
	)
	p.API.LogInfo(
		"##################################",
	)

	return resp, nil
}
