package bot

import (
	"fmt"
	"sort"

	"github.com/CarlFlo/GoDiscordBotTemplate/bot/structs"

	"github.com/bwmarrin/discordgo"
)

var helpString string
var helpStringAdmin string

func generateHelp() {

	helpUsr := []string{}
	helpAdmin := []string{}

	helpString = "[User]\n"
	helpStringAdmin = "\n[Admin]\n"

	for cmd, data := range validCommands {

		helpSyntax := ""
		if len(data.helpSyntax) > 0 {
			helpSyntax = fmt.Sprintf(" %s", data.helpSyntax)
		}

		if data.requiredPermission == enumAdmin {
			helpAdmin = append(helpAdmin, cmd+helpSyntax)
		} else {
			helpUsr = append(helpUsr, cmd+helpSyntax)
		}
	}

	// Sort the commands
	sort.Strings(helpUsr)
	sort.Strings(helpAdmin)

	// Adds to string
	for _, str := range helpUsr {
		helpString += str + ", "
	}
	for _, str := range helpAdmin {
		helpStringAdmin += str + ", "
	}

	// Trims the end
	helpString = helpString[:len(helpString)-2]
	helpStringAdmin = helpStringAdmin[:len(helpStringAdmin)-2]

}

/*
 Automatically generate help
*/
func help(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput) {

	start := "```ini\n"
	end := "\n[Note]\nCommands are not case sensitive.\n```"

	// Admins will get additional help
	if input.IsAdmin() {
		s.ChannelMessageSend(m.ChannelID, start+helpString+helpStringAdmin+end)
	} else {
		s.ChannelMessageSend(m.ChannelID, start+helpString+end)
	}

}
