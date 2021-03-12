package commands

import (
	"log"

	"github.com/CarlFlo/GoDiscordBotTemplate/bot/commands/cmdutils"
	"github.com/CarlFlo/GoDiscordBotTemplate/bot/structs"
	"github.com/CarlFlo/GoDiscordBotTemplate/config"

	"github.com/bwmarrin/discordgo"
)

// Reload - Reloads the configuration without restarting the application
func Reload(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput) {

	if err := config.ReloadConfig(); err != nil {
		log.Println("Could not reload config!")
		return
	}

	cmdutils.PipeDirectMessage(s, m, "Config reloaded")
}
