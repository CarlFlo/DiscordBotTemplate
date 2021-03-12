package commands

import (
	"fmt"
	"log"

	"main/bot/commands/cmdutils"
	"main/bot/structs"
	"main/config"

	"github.com/bwmarrin/discordgo"
)

// BotInvite - Sends back the invite link to the bot
func BotInvite(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput) {

	if len(config.CONFIG.BotInfo.ClientID) == 0 {
		log.Printf("ClientID not set in config file")
		cmdutils.PipeDirectMessage(s, m, "Unable to create bot invite. Contact the administrator")
		return
	}

	inviteLink := fmt.Sprintf("https://discordapp.com/oauth2/authorize?&client_id=%s&scope=bot&permissions=%d", config.CONFIG.BotInfo.ClientID, config.CONFIG.BotInfo.Permission)

	cmdutils.PipeDirectMessage(s, m, inviteLink)
}
