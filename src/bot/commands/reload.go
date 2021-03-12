package commands

import (
	"log"

	"main/bot/commands/cmdutils"
	"main/bot/structs"
	"main/config"

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
