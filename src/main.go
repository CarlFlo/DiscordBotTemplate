package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"main/bot"
	"main/bot/commands/cmdutils"
	"main/config"
	"main/utils"
)

// https://discordapp.com/oauth2/authorize?&client_id=643191140849549333&scope=bot&permissions=37211200

func init() {
	utils.Clear()

	rand.Seed(time.Now().UTC().UnixNano())

	log.Printf("Running on %s\n", runtime.GOOS)

	if err := config.LoadConfiguration(); err != nil {
		panic(err)
	}

	// Loads all the valid commands into a map
	bot.MapValidCommands()

	log.Printf("Version %s\n", config.CONFIG.Version)

	go cmdutils.MessagePipeLoop()
}

func main() {

	if err := bot.StartBot(); err != nil {
		log.Fatalln(err)
		return
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
