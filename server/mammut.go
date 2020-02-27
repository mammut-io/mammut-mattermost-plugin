package main

import (
	"encoding/json"
)

//MammutResponse response from mammut api
type MammutResponse struct {
	UserID    string
	ChannelID string
	Message   string
}

//MammutUserPayload response from mammut api
type MammutUserPayload struct {
	UserType         string
	Username         string
	MattermostUserID string
}

//MammutPayloadToJSON marshals MammutUserPayload
func (p *Plugin) MammutPayloadToJSON(mammutuserpayload *MammutUserPayload) ([]byte, error) {
	var jsonLoad map[string]string
	jsonLoad = make(map[string]string)
	jsonLoad["user-type"] = mammutuserpayload.UserType
	jsonLoad["name"] = mammutuserpayload.Username
	jsonLoad["mattermost-user-id"] = mammutuserpayload.MattermostUserID
	p.API.LogInfo(
		">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>",
	)
	p.API.LogInfo(
		"jsonload",
		"jsonload", jsonLoad,
	)
	jsonLoadResult, err := json.Marshal(jsonLoad)
	p.API.LogInfo(
		"jsonload marshaled",
		"jsonload marshaled", jsonLoadResult,
	)
	jsonResult, err := json.Marshal(mammutuserpayload)
	p.API.LogInfo(
		"jsonResult",
		"jsonResult", jsonResult,
	)
	if err != nil {
		return nil, err
	}
	return jsonResult, err
}
