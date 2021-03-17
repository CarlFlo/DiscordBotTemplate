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

	log.Printf("Version %s\n", config.CONFIG.Version)
}

func main() {

	session := bot.StartBot()

	time.Sleep(500 * time.Millisecond) // Added this sleep so the message under will come last
	// Keeps bot from closing. Waits for CTRL-C
	log.Printf("Press CTRL-C to initiate shutdown\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	log.Printf("Shutting down!\n")

	// Run cleanup code here
	session.Close() // Stops the discord bot
}
