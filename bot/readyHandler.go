package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func readyHandler(s *discordgo.Session, ready *discordgo.Ready) {

	statusMessage := fmt.Sprintf("on %d servers", len(s.State.Guilds))

	s.UpdateStreamingStatus(0, statusMessage, "https://www.youtube.com/watch?v=3glxLWVkbSs")

	//s.UpdateGameStatus(0, statusMessage)

	log.Printf("Ready with %d guilds\n", len(s.State.Guilds))
}
