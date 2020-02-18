package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v5/model"
)

func (p *Plugin) postMammutPluginMessage(id, msg string) *model.AppError {
	configuration := p.getConfiguration()
	p.API.LogInfo(
		"##################################",
	)
	p.API.LogInfo(
		"id",
		"id", id,
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
