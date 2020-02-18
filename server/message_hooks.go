package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// MessageHasBeenPosted is invoked after the message has been committed to the database. If you
// need to modify or reject the post, see MessageWillBePosted Note that this method will be called
// for posts created by plugins, including the plugin that created the post.
//
// This demo implementation logs a message to the demo channel whenever a message is posted,
// unless by the demo plugin user itself.
func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	configuration := p.getConfiguration()

	if configuration.disabled {
		return
	}

	// Ignore posts by the demo plugin user and demo plugin bot.
	if post.UserId == p.botID || post.UserId == configuration.demoUserID {
		return
	}

	user, err := p.API.GetUser(post.UserId)
	if err != nil {
		p.API.LogError("failed to query user", "user_id", post.UserId)
		return
	}

	channel, err := p.API.GetChannel(post.ChannelId)
	if err != nil {
		p.API.LogError("failed to query channel", "channel_id", post.ChannelId)
		return
	}

	msg := fmt.Sprintf("MessageHasBeenPosted: @%s, ~%s", user.Username, channel.Name)
	p.API.LogDebug(
		"chanel id before post",
		"chanelid", channel.TeamId,
	)
	if err := p.postPluginMessage(channel.TeamId, msg); err != nil {
		p.API.LogError(
			"failed to post MessageHasBeenPosted message",
			"channel_id", channel.Id,
			"user_id", user.Id,
			"error", err.Error(),
		)
	}

	// Check if the Random Secret was posted
	if strings.Contains(post.Message, configuration.RandomSecret) {
		msg = fmt.Sprintf("The random secret %q has been entered by @%s!\n%s",
			configuration.RandomSecret, user.Username, configuration.SecretMessage,
		)
		if err := p.postPluginMessage(channel.TeamId, msg); err != nil {
			p.API.LogError(
				"failed to post random secret message",
				"channel_id", channel.Id,
				"user_id", user.Id,
				"error", err.Error(),
			)
		}
	}

	if strings.Contains(post.Message, strconv.Itoa(configuration.SecretNumber)) {
		msg = fmt.Sprintf("The random number %d has been entered by @%s!",
			configuration.SecretNumber, user.Username)
		if err := p.postPluginMessage(channel.TeamId, msg); err != nil {
			p.API.LogError(
				"failed to post random secret message",
				"channel_id", channel.Id,
				"user_id", user.Id,
				"error", err.Error(),
			)
		}
	}
}
