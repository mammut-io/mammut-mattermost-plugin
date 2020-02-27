package main

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

//MammutPayloadToMAP marshals MammutUserPayload
func (p *Plugin) MammutPayloadToMAP(mammutuserpayload *MammutUserPayload) (map[string]string, error) {
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
	return jsonLoad, nil
}
