package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

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
	testbody := &MammutResponse{
		UserID:   p.botID,
		ChanelID: "wnkycixbb3bgjghobwb99ndjka",
		Message:  "Pereira",
	}

	jsonTest, err := json.Marshal(testbody)
	req, err := http.NewRequest("POST", "http://localhost:8065/plugins/com.mattermost.mammut-mattermos-plugin/mammuthook", strings.NewReader(string(jsonTest)))
	if err != nil {
		return nil, model.NewAppError("DoActionRequest1", "api.post.do_action.action_integration.app_error", nil, err.Error(), http.StatusBadRequest)
	}
	req.Header.Add("Mattermost-User-Id", "theuserid")
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
