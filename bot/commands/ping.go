package commands

import (
	"github.com/CarlFlo/GoDiscordBotTemplate/bot/commands/cmdutils"
	"github.com/CarlFlo/GoDiscordBotTemplate/bot/structs"

	"github.com/bwmarrin/discordgo"
)

// Ping - Sends back a 'Pong' message
func Ping(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput) {

	cmdutils.SendDirectMessage(s, m, "Pong!")
}
