package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Redo this so this isnt required https://www.youtube.com/watch?v=y_eIBmt3JdY

// CONFIG holds all the config data
var CONFIG *configStruct

type configStruct struct {
	Token                string            `json:"token"`
	BotPrefix            string            `json:"botPrefix"`
	Version              string            `json:"version"`
	OwnerID              string            `json:"ownerID"`
	DispConfOnStart      bool              `json:"dispConfOnStart"`
	BoundChannels        []string          `json:"boundChannels"`
	AllowDirectMessages  bool              `json:"allowDirectMessages"`
	MessageHandlerBuffer int               `json:"messageHandlerBuffer"`
	BotInfo              botInfo           `json:"botInfo"`
	MessageProcessing    messageProcessing `json:"messageProcessing"`
}

type botInfo struct {
	ClientID   string `json:"clientID"`
	Permission uint64 `json:"permission"`
}

type messageProcessing struct {
	MessageLengthLimit    int `json:"messageLengthLimit"`
	MaxIncommingMsgLength int `json:"maxIncommingMsgLength"`
}

// ReloadConfig is a wrapper function for reloading the config. For clarity
func ReloadConfig() error {
	return readConfig()
}

// readConfig will read the config file
func readConfig() error {
	//log.Println("Reading from config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return err
	}

	if err = json.Unmarshal(file, &CONFIG); err != nil {
		return err
	}

	if CONFIG.DispConfOnStart {
		log.Printf("Config:\n%s\n", string(file))
	}

	return nil
}

// createConfig creates the default config file
func createConfig() error {

	// Default
	configStruct := configStruct{
		Token:                "<TokenHere>",
		BotPrefix:            ",",
		Version:              "2021-03-12",
		OwnerID:              "<YourDiscordIDHere>",
		DispConfOnStart:      false,
		BoundChannels:        []string{},
		AllowDirectMessages:  true,
		MessageHandlerBuffer: 5,
		BotInfo: botInfo{
			ClientID:   "<BotIDHere>",
			Permission: 37211200,
		},
		MessageProcessing: messageProcessing{
			MessageLengthLimit:    1850,
			MaxIncommingMsgLength: 100,
		},
	}

	jsonData, _ := json.MarshalIndent(configStruct, "", "   ")
	err := ioutil.WriteFile("config.json", jsonData, 0644)

	return err
}

// LoadConfiguration loads the configuration file into memory
func LoadConfiguration() error {
	err := readConfig()
	if err != nil {
		if err = createConfig(); err != nil {
			return err
		}
		if err = readConfig(); err != nil {
			return err
		}
	}
	return nil
}
