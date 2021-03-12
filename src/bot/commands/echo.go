package commands

import (
	"strings"

	"github.com/CarlFlo/GoDiscordBotTemplate/bot/commands/cmdutils"
	"github.com/CarlFlo/GoDiscordBotTemplate/bot/structs"

	"github.com/bwmarrin/discordgo"
)

// Echo - echoes the message
func Echo(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput) {

	output := strings.Join(input.GetArgs(), " ")

	cmdutils.PipeChannelMessage(s, m.ChannelID, output)
}
