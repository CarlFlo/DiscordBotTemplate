package commands

import (
	"main/bot/commands/cmdutils"
	"main/bot/structs"

	"github.com/bwmarrin/discordgo"
)

// Ping - Sends back a 'Pong' message
func Ping(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput) {

	cmdutils.PipeDirectMessage(s, m, "Pong!")
}
