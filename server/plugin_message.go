package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

func (p *Plugin) postMammutPluginMessageToAPI(channelID string, user *model.User, msg string) *model.AppError {
	configuration := p.getConfiguration()

	if configuration.disabled {
		return nil
	}
	configMattermostURL := p.API.GetConfig().ServiceSettings.SiteURL
	//TODO: this is beacuse mammut not suporting ":" yet
	configMattermostCleanURL := strings.Replace(*configMattermostURL, "https://", "", -1)
	mammutMessageBody := &MammutMessage{
		UserID:     user.Id,
		BotID:      p.botID,
		ChannelID:  channelID,
		Message:    msg,
		UserEmail:  user.Email,
		DomainName: configMattermostCleanURL,
	}
	p.API.LogInfo(
		"##################################",
	)
	p.API.LogInfo(
		"mammutMessageBody",
		"mammutMessageBody", mammutMessageBody,
	)
	jsonBody, jsonErr := json.Marshal(mammutMessageBody)
	//TODO:no se si esta es la forma correcta de manejar el error
	if jsonErr != nil {
		return model.NewAppError("postMammutPluginMessageToAPI", "plugin.MessageHasBeenPosted.postMammutPluginMessageToAPI.json.marshal", nil, jsonErr.Error(), http.StatusBadRequest)
	}
	url := configuration.MammutAPIURL + "/channel/mattermost"
	_, err := p.doActionRequest(url, jsonBody)

	return err
}
