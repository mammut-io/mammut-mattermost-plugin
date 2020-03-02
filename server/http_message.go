package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

// Perform an HTTP POST request to an integration's action endpoint.
// Caller must consume and close returned http.Response as necessary.
// For internal requests, requests are routed directly to a plugin ServerHTTP hook
func (p *Plugin) doActionRequest(rawURL string, body []byte) (*http.Response, *model.AppError) {
	p.API.LogInfo(
		"##################################",
	)
	p.API.LogInfo(
		"rawURLPath",
		"rawURL", rawURL,
	)
	req, err := http.NewRequest("POST", rawURL, strings.NewReader(string(body)))
	if err != nil {
		return nil, model.NewAppError("DoActionRequest1", "api.post.do_action.action_integration.app_error", nil, err.Error(), http.StatusBadRequest)
	}
	req.Header.Add("Mattermost-User-Id", p.botID)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Allow access to plugin routes for action buttons
	httpClient := &http.Client{}
	resp, httpErr := httpClient.Do(req)
	if httpErr != nil {
		return nil, model.NewAppError("DoActionRequest2", "api.post.do_action.action_integration.app_error", nil, "err="+httpErr.Error(), http.StatusBadRequest)
	}

	p.API.LogInfo(
		"response",
		"resp.StatusCode", resp.StatusCode,
	)
	p.API.LogInfo(
		"##################################",
	)
	if resp.StatusCode != http.StatusOK {
		return resp, model.NewAppError("DoActionRequest3", "api.post.do_action.action_integration.app_error", nil, fmt.Sprintf("status=%v", resp.StatusCode), http.StatusBadRequest)
	}

	return resp, nil
}
