package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func presenceUpdateHandler(s *discordgo.Session, m *discordgo.PresenceUpdate) {

	log.Printf("Username: %s Status: %s Email: %s\n", m.User.Username, m.Status, m.User.Email)

}
