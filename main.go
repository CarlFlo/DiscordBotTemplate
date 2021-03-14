package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/CarlFlo/GoDiscordBotTemplate/bot"
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
	if len(config.CONFIG.Token) == 0 {
		log.Fatalln("No token provided, check the configuration file")
	}

	// Loads all the valid commands into a map
	bot.MapValidCommands()

	log.Printf("Version %s\n", config.CONFIG.Version)
}

func main() {

	session := bot.StartBot()
	log.Println("Bot is connected!")

	// Keeps bot from closing. Waits for CTRL-C
	log.Printf("Now running. Press CTRL-C to exit\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Stops the bot
	session.Close()
}
