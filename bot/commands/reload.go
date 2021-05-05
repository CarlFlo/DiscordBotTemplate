package commands

import (
	"github.com/CarlFlo/bord"
	"github.com/CarlFlo/discordBotTemplate/bot/commands/cmdutils"
	"github.com/CarlFlo/discordBotTemplate/bot/structs"
	"github.com/CarlFlo/discordBotTemplate/config"

	"github.com/bwmarrin/discordgo"
)

// Reload - Reloads the configuration without restarting the application
func Reload(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput) {

	if err := config.ReloadConfig(); err != nil {
		bord.Error("Could not reload config! %s", err)
		return
	}

	cmdutils.SendDirectMessage(s, m, "Config reloaded")
}
