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
	"github.com/CarlFlo/bord"
)

// https://discordapp.com/oauth2/authorize?&client_id=643191140849549333&scope=bot&permissions=37211200

func init() {

	utils.Clear()
	rand.Seed(time.Now().UTC().UnixNano())

	bord.Debug("Running on %s", runtime.GOOS)

	if err := config.LoadConfiguration(); err != nil {
		log.Fatalln(err)
	}

	bord.Debug("Version %s", config.CONFIG.Version)
}

func main() {

	session := bot.StartBot()

	time.Sleep(500 * time.Millisecond) // Added this sleep so the message under will come last
	// Keeps bot from closing. Waits for CTRL-C
	bord.Info("Press CTRL-C to initiate shutdown")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	bord.Info("Shutting down!")

	// Run cleanup code here
	close(sc)
	session.Close() // Stops the discord bot
}
