package cmdutils

import (
	"log"

	"github.com/CarlFlo/GoDiscordBotTemplate/config"

	"github.com/bwmarrin/discordgo"
)

var pipe chan msg

// Inspiration https://stackoverflow.com/questions/35926173/implementing-promise-with-channels-in-go
type future struct {
	result *discordgo.Message
	err    error
}

type msg struct {
	s         *discordgo.Session
	channelID string
	content   string
	f         chan future
}

// PipeDirectMessage - Sends a direct messag to a user
func PipeDirectMessage(s *discordgo.Session, m *discordgo.MessageCreate, content string) (*discordgo.Message, error) {
	ch, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return PipeChannelMessage(s, ch.ID, content)
}

// PipeChannelMessage -
func PipeChannelMessage(s *discordgo.Session, channelID, content string) (*discordgo.Message, error) {

	future := make(chan future, 1)

	pipe <- msg{s, channelID, content, future}
	fBack := <-future
	return fBack.result, fBack.err
}

// MessagePipeLoop all messages goes through this function
func MessagePipeLoop() {

	// Initialise channel
	pipe = make(chan msg, config.CONFIG.MessageHandlerBuffer)

	for pipeMsg := range pipe {

		discordMsg, err := pipeMsg.s.ChannelMessageSend(pipeMsg.channelID, pipeMsg.content)
		pipeMsg.f <- future{discordMsg, err}
	}
}
