package bot

import (
	"main/config"
)

func isOwner(discordID string) bool {
	return discordID == config.CONFIG.OwnerID
}
