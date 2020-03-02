package main

import (
	"encoding/json"
)

//MammutMessage the struct needed to send mesage from mattermost to mammutAPI
type MammutMessage struct {
	UserID     string `json:"user_id"`
	BotID      string `json:"bot_id"`
	ChannelID  string `json:"channel_id"`
	Message    string `json:"message"`
	UserEmail  string `json:"user_email"`
	DomainName string `json:"domain_name"`
}

//MammutResponse response from mammut api
type MammutResponse struct {
	UserID    string `json:"user_id"`
	ChannelID string `json:"channel_id"`
	Message   string `json:"message"`
}

//MammutUserPayload response from mammut api
type MammutUserPayload struct {
	UserType         string   `json:"user-type"`
	MainEmail        string   `json:"main-email"`
	Username         string   `json:"name"`
	MattermostUserID []string `json:"mattermost-user-id"`
}

//TaskResultBasic used by MammutUserCreationResponse for maping reponse json on create user resquest
type TaskResultBasic struct {
	AffectedElementID   int64         `json:"affectedElementId"`
	AffectedElementName string        `json:"affectedElementName"`
	AffectedElementType string        `json:"affectedElementType"`
	TaskIDList          []interface{} `json:"taskIdList"`
}

//MammutUserCreationResponse is to mapo the reponse of ammut creation from json on create user resquest
type MammutUserCreationResponse struct {
	Status     string            `json:"status"`
	Taskresult []TaskResultBasic `json:"taskresult"`
}

//MammutPayloadToJSON marshals MammutUserPayload
func (p *Plugin) MammutPayloadToJSON(mammutuserpayload *MammutUserPayload) ([]byte, error) {
	jsonLoadResult, err := json.Marshal(mammutuserpayload)
	if err != nil {
		return nil, err
	}

	return jsonLoadResult, nil
}
