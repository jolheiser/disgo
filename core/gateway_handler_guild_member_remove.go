package core

import (
	"github.com/DisgoOrg/disgo/discord"
)

type guildMemberRemoveData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	User    discord.User      `json:"user"`
}

// GuildMemberRemoveHandler handles core.GuildMemberRemoveGatewayEvent
type GuildMemberRemoveHandler struct{}

// EventType returns the core.GatewayGatewayEventType
func (h *GuildMemberRemoveHandler) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeGuildMemberRemove
}

// New constructs a new payload receiver for the raw gateway event
func (h *GuildMemberRemoveHandler) New() interface{} {
	return &guildMemberRemoveData{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *GuildMemberRemoveHandler) HandleGatewayEvent(bot *Bot, sequenceNumber int, v interface{}) {
	memberData := *v.(*guildMemberRemoveData)

	bot.EntityBuilder.CreateUser(memberData.User, CacheStrategyYes)

	member := bot.Caches.MemberCache().GetCopy(memberData.GuildID, memberData.User.ID)

	bot.Caches.MemberCache().Remove(memberData.GuildID, memberData.User.ID)

	bot.EventManager.Dispatch(&GuildMemberLeaveEvent{
		GenericGuildMemberEvent: &GenericGuildMemberEvent{
			GenericGuildEvent: &GenericGuildEvent{
				GenericEvent: NewGenericEvent(bot, sequenceNumber),
				Guild:        bot.Caches.GuildCache().Get(memberData.GuildID),
			},
			Member: member,
		},
	})
}
