package commands

import (
	"github.com/CarlFlo/GoDiscordBotTemplate/bot/commands/cmdutils"
	"github.com/CarlFlo/GoDiscordBotTemplate/bot/structs"

	"github.com/bwmarrin/discordgo"
)

// Ping - Sends back a 'Pong' message
func Ping(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput) {

	cmdutils.SendDirectMessage(s, m, "Pong!")

	/*
		// Send ping
		pingMsg, err := cmdutils.SendDirectMessage(s, m, "Pinging...")
		if err != nil {
			log.Printf("Error: %s", err)
			return
		}

		// Update message
		s.ChannelMessageEdit(pingMsg.ChannelID, pingMsg.ID, "Pinging... :bar_chart:")

		// Save this and pass it so the messageUpdateHandler chan access it
		// pingMsg.ID
		// discordgo.SnowflakeTimestamp(pingMsg.ID)

		// elapsed := time.Since(initalTime)
		// This info will we used to calculate the ping
		// Send the ID, match the ID and use the cached timestamp to compare
	*/
}
