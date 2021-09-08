package core

import (
	"github.com/DisgoOrg/disgo/discord"
)

type webhooksUpdateData struct {
	GuildID   discord.Snowflake `json:"guild_id"`
	ChannelID discord.Snowflake `json:"channel_id"`
}

// WebhooksUpdateHandler handles core.GatewayEventWebhooksUpdate
type WebhooksUpdateHandler struct{}

// EventType returns the raw core.GatewayGatewayEventType
func (h *WebhooksUpdateHandler) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeWebhooksUpdate
}

// New constructs a new payload receiver for the raw gateway event
func (h *WebhooksUpdateHandler) New() interface{} {
	return &webhooksUpdateData{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *WebhooksUpdateHandler) HandleGatewayEvent(bot *Bot, sequenceNumber int, v interface{}) {

}
