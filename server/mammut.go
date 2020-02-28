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
	UserType         string `json:"user-type"`
	Username         string `json:"name"`
	MattermostUserID string `json:"mattermost-user-id"`
}

//TaskResultBasic used by MammutUserCreationResponse for maping reponse json
type TaskResultBasic struct {
	AffectedElementID   int64         `json:"affectedElementId"`
	AffectedElementName string        `json:"affectedElementName"`
	AffectedElementType string        `json:"affectedElementType"`
	TaskIDList          []interface{} `json:"taskIdList"`
}

//MammutUserCreationResponse is to mapo the reponse of ammut creation from json
type MammutUserCreationResponse struct {
	Status     string            `json:"status"`
	Taskresult []TaskResultBasic `json:"taskresult"`
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
	if err != nil {
		return nil, err
	}

	return jsonLoadResult, nil
}
