package main

import (
	"fmt"

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
	userbot, err := p.API.GetUser(p.botID)
	p.API.LogInfo(
		"##################################",
	)
	p.API.LogInfo(
		"bot",
		"userbot", userbot,
	)
	//TODO tengo que hacer esto o con el post.ChannelId basta?
	channel, err := p.API.GetChannel(post.ChannelId)
	if err != nil {
		p.API.LogError("failed to query channel", "channel_id", post.ChannelId)
		return
	}

	directChannel, err := p.API.GetDirectChannel(post.UserId, p.botID)
	if err != nil {
		p.API.LogError(
			"Failed to get channel for client and bot",
			"post_userId", post.UserId,
			"bot_demo_user", p.botID,
			"error", err.Error(),
		)
		return
	}
	p.API.LogInfo(
		"chanels before comparison",
		"channel_id", channel.Id,
		"direct_channel_id", directChannel.Id,
	)
	msg := post.Message
	if directChannel.Id == channel.Id {
		if err := p.postMammutPluginMessageToAPI(channel.Id, msg); err != nil {
			p.API.LogError(
				"failed to post MessageHasBeenPosted message to mammut API",
				"channel_id", channel.Id,
				"direct_channel_id", directChannel.Id,
				"error", err.Error(),
			)
		}
	}
	//TODO: esto no va aqui, deberia ser on http event
	msgResp := fmt.Sprintf("MessageHasBeenPosted: @%s, ~%s", user.Username, channel.Name)
	if directChannel.Id == channel.Id {
		if err := p.postMammutPluginMessage(channel.Id, msgResp); err != nil {
			p.API.LogError(
				"failed to post MessageHasBeenPosted message",
				"channel_id", channel.Id,
				"direct_channel_id", directChannel.Id,
				"error", err.Error(),
			)
		}
	}
}
