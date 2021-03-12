package bot

import (
	"fmt"
	"unicode"

	"main/bot/commands"
	"main/bot/structs"

	"github.com/bwmarrin/discordgo"
)

const (
	enumUser uint8 = iota
	enumAdmin
)

// Command type for sorting similar commands together
const (
	typeUser uint8 = iota
	typeMisc
	typeAdmin
)

type command struct {
	function           func(s *discordgo.Session, m *discordgo.MessageCreate, input structs.CmdInput)
	requiredPermission uint8
	helpSyntax         string
	commandType        uint8
}

// Add variable to specify that command only can be run in a guild, not in a directmessage

var validCommands = make(map[string]command)

// MapValidCommands will initialize a map
// with all the valid functions that can be run
func MapValidCommands() {

	/* all keys MUST be lowercase */
	// Admin commands
	validCommands["reload"] = command{
		function:           commands.Reload,
		requiredPermission: enumAdmin,
		commandType:        typeAdmin}

	validCommands["presence"] = command{
		function:           commands.Presence,
		requiredPermission: enumAdmin,
		helpSyntax:         "[v verbose, d dump]",
		commandType:        typeAdmin}

	// Perm User - User commands
	validCommands["help"] = command{
		function:           commands.Help,
		requiredPermission: enumUser,
		commandType:        typeUser}

	validCommands["h"] = command{
		function:           help,
		requiredPermission: enumUser,
		commandType:        typeUser}

	// Perm User - Misc commands
	validCommands["ping"] = command{
		function:           commands.Ping,
		requiredPermission: enumUser,
		commandType:        typeMisc}

	validCommands["botinvite"] = command{
		function:           commands.BotInvite,
		requiredPermission: enumUser,
		commandType:        typeMisc}

	validCommands["echo"] = command{
		function:           commands.Echo,
		requiredPermission: enumUser,
		helpSyntax:         "<text>",
		commandType:        typeMisc}

	// Validates the keys so no-one is uppercase
	validateKeys()
	generateHelp()
}

// Validates that all the keys are lowercase
func validateKeys() {
	for key := range validCommands {
		for _, char := range key {
			if !unicode.IsLower(char) {
				panic(fmt.Sprintf("key: '%s' contains one or more non lowercase characters => '%c'", key, char))
			}
		}
	}
}
