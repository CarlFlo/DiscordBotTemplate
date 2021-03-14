package bot

import (
	"log"

	"github.com/CarlFlo/GoDiscordBotTemplate/config"

	"github.com/bwmarrin/discordgo"
)

// StartBot starts the bot and returns any errors that might occu
func StartBot() *discordgo.Session {

	session, err := discordgo.New("Bot " + config.CONFIG.Token) // Creates bot
	if err != nil {
		return nil
	}

	// Adds message handler (https://github.com/bwmarrin/discordgo/blob/37088aefec2241139e59b9b804f193b539be25d6/eventhandlers.go#L937)
	session.AddHandler(messageHandler)
	session.AddHandler(presenceUpdateHandler)

	// Attempts to open connection
	err = session.Open()
	if err != nil {
		log.Fatalln(err)
	}

	// Returns session
	return session
}
