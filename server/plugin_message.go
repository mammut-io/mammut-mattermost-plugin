package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/v5/model"
)

func (p *Plugin) postPluginMessage(id, msg string) *model.AppError {
	configuration := p.getConfiguration()
	p.API.LogDebug(
		"##################################",
	)
	p.API.LogDebug(
		"configuration.demoChannelIDs[id]",
		"configuration chanel id", configuration.demoChannelIDs[id],
	)
	p.API.LogDebug(
		"id",
		"id", id,
	)
	p.API.LogDebug(
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
		ChannelId: configuration.demoChannelIDs[id],
		Message:   msg,
	})

	return err
}
