package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
)

func (p *Plugin) postMammutPluginMessageToAPI(channelID, msg string) *model.AppError {
	configuration := p.getConfiguration()

	if configuration.disabled {
		return nil
	}

	body := &MammutResponse{
		UserID:    p.botID,
		ChannelID: channelID,
		Message:   msg,
	}
	jsonBody, jsonErr := json.Marshal(body)
	//TODO:no se si esta es la forma correcta de manejar el error
	if jsonErr != nil {
		return model.NewAppError("postMammutPluginMessageToAPI", "plugin.MessageHasBeenPosted.postMammutPluginMessageToAPI.json.marshal", nil, jsonErr.Error(), http.StatusBadRequest)
	}
	_, err := p.doActionRequest(configuration.MammutAPIURL, jsonBody)

	return err
}
func (p *Plugin) postMammutPluginMessage(id, msg string) *model.AppError {
	configuration := p.getConfiguration()
	p.API.LogInfo(
		"##################################",
	)
	p.API.LogInfo(
		"id",
		"channelid", id,
		"botid", p.botID,
	)
	p.API.LogInfo(
		"##################################",
	)

	if configuration.disabled {
		return nil
	}

	if configuration.EnableMentionUser {
		msg = fmt.Sprintf("tag @%s | %s", configuration.MentionUser, msg)
	}
	msg = fmt.Sprintf("%s%s%s", configuration.TextStyle, msg, configuration.TextStyle)

	_, err := p.API.CreatePost(&model.Post{
		UserId:    p.botID,
		ChannelId: id,
		Message:   msg,
	})

	return err
}
