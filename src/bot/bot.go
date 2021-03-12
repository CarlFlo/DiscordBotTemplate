package bot

import (
	"main/config"

	"github.com/bwmarrin/discordgo"
)

// StartBot starts the bot and returns any errors that might occu
func StartBot() error {

	bot, err := discordgo.New("Bot " + config.CONFIG.Token) // Creates bot
	if err != nil {
		return err
	}

	// Adds message handler (https://github.com/bwmarrin/discordgo/blob/37088aefec2241139e59b9b804f193b539be25d6/eventhandlers.go#L937)
	bot.AddHandler(messageHandler)
	bot.AddHandler(presenceUpdateHandler)

	// Opens connection
	return bot.Open()
}
