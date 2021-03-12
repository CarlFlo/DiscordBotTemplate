package commands

import (
	"main/bot/commands/cmdutils"
	"main/bot/structs"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Echo - echoes the message
func Echo(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput) {

	output := strings.Join(input.GetArgs(), " ")

	cmdutils.PipeChannelMessage(s, m.ChannelID, output)
}
