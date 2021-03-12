package bot

import (
	"github.com/CarlFlo/GoDiscordBotTemplate/config"
)

func isOwner(discordID string) bool {
	return discordID == config.CONFIG.OwnerID
}
