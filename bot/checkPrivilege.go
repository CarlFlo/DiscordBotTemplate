package bot

import (
	"github.com/CarlFlo/DiscordBotTemplate/config"
)

func isOwner(discordID string) bool {
	return discordID == config.CONFIG.OwnerID
}
