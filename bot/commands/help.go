package commands

import (
	"github.com/CarlFlo/GoDiscordBotTemplate/bot/structs"

	"github.com/bwmarrin/discordgo"
)

// Help - Displays the help message
// Admins will get additional help
func Help(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput) {

	/*
		TODO: Generate help automatically.
		From validCommands map. It will need a helpType if its a [User], or [Misc]
	*/
	var output string

	appendUserHelp(&output)

	if input.IsAdmin() {
		appendAdminHelp(&output)
	}

	start := "```ini\n"
	end := "[Note]\nCommands are not case sensitive.\n```"

	s.ChannelMessageSend(m.ChannelID, start+output+end)
}

func appendUserHelp(output *string) {

	*output = `
[User]
help
[Misc]
ping, botinvite, echo <text>
`
}

func appendAdminHelp(output *string) {

	*output += `[Admin]
reload, presence [v verbose, d dump]
`
}
