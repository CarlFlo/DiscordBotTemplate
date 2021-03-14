package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/CarlFlo/GoDiscordBotTemplate/bot"
	"github.com/CarlFlo/GoDiscordBotTemplate/bot/commands/cmdutils"
	"github.com/CarlFlo/GoDiscordBotTemplate/config"
	"github.com/CarlFlo/GoDiscordBotTemplate/utils"
)

// https://discordapp.com/oauth2/authorize?&client_id=643191140849549333&scope=bot&permissions=37211200

func init() {
	utils.Clear()

	rand.Seed(time.Now().UTC().UnixNano())

	log.Printf("Running on %s\n", runtime.GOOS)

	if err := config.LoadConfiguration(); err != nil {
		log.Fatalln(err)
	}

	// Checks if bot token is present
	if config.CONFIG.Token == "TokenHere" {
		log.Fatalln("No token provided, check the configuration file")
	}

	// Loads all the valid commands into a map
	bot.MapValidCommands()

	log.Printf("Version %s\n", config.CONFIG.Version)

	go cmdutils.MessagePipeLoop()
}

func main() {

	if err := bot.StartBot(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Bot is connected!")

	// Keeps bot from closing
	stayOpen()
}

func stayOpen() {

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
